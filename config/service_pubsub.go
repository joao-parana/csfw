// Copyright 2015-2016, Cyrill @ Schumacher.fm and the CoreStore contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"errors"
	"sync"

	"github.com/corestoreio/csfw/store/scope"
	"github.com/juju/errgo"
)

// ErrPublisherClosed will returned when the channel has been closed.
var ErrPublisherClosed = errors.New("config Service Publisher already closed")

// MessageReceiver allows you to listen to write actions. The order of calling
// each subscriber is totally random. If a subscriber panics, it gets securely
// removed without crashing the whole system.
// This interface should be implemented in other packages.
// The Subscriber interface requires the MessageReceiver interface.
type MessageReceiver interface {
	// MessageConfig when a configuration value will be written this function
	// gets called to allow you to listen to changes. Path is never empty.
	// Path may contains up to three levels. For more details see the Subscriber
	// interface of this package. If an error will be returned, the subscriber
	// gets unsubscribed/removed.
	MessageConfig(path string, sg scope.Scope, id int64) error
}

// Subscriber represents the overall service to receive subscriptions from
// MessageReceiver interfaces. This interface is at the moment only implemented
// by the config.Service.
type Subscriber interface {
	// Subscribe subscribes a MessageReceiver to a path.
	// Path allows you to filter to which path or part of a path you would like to listen.
	// A path can be e.g. "system/smtp/host" to receive messages by single host changes or
	// "system/smtp" to receive message from all smtp changes or "system" to receive changes
	// for all paths beginning with "system". A path is equal to a topic in a PubSub system.
	// Path cannot be empty means you cannot listen to all changes.
	// Returns a unique identifier for the Subscriber for later removal, or an error.
	Subscribe(path string, s MessageReceiver) (subscriptionID int, err error)
}

// pubSub embedded pointer struct into the Service
type pubSub struct {
	// subMap, subscribed writers are getting called when a write event
	// will happen. String is the path (aka topic) and int the Subscriber ID for later
	// removal.
	subMap     map[string]map[int]MessageReceiver
	subAutoInc int // subAutoInc increased whenever a Subscriber has been added
	mu         sync.RWMutex
	publishArg chan arg
	stop       chan struct{} // terminates the goroutine
	closeErr   chan error    // this one tells us that the go routine has really been terminated
	closed     bool          // if Close() has been called the config.Service can still Write() without panic
}

// Close closes the internal channel for the pubsub Goroutine. Prevents a leaking
// Goroutine.
func (s *pubSub) Close() error {
	if s.closed {
		return ErrPublisherClosed
	}
	defer func() { close(s.closeErr) }() // last close(s.closeErr) does not work and panics
	s.closed = true
	s.stop <- struct{}{}
	close(s.publishArg)
	close(s.stop)
	//close(s.closeErr)
	return <-s.closeErr
}

// Subscribe adds a Subscriber to be called when a write event happens.
// See interface Subscriber for a detailed description.
func (s *pubSub) Subscribe(path string, mr MessageReceiver) (subscriptionID int, err error) {
	if path == "" {
		return 0, ErrPathEmpty
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.subAutoInc++
	subscriptionID = s.subAutoInc

	if _, ok := s.subMap[path]; !ok {
		s.subMap[path] = make(map[int]MessageReceiver)
	}
	s.subMap[path][subscriptionID] = mr

	return
}

// Unsubscribe removes a subscriber with a specific ID.
func (s *pubSub) Unsubscribe(subscriptionID int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for path, subs := range s.subMap {
		if _, ok := subs[subscriptionID]; ok {
			delete(s.subMap[path], subscriptionID) // mem leaks?
			if len(s.subMap[path]) == 0 {
				delete(s.subMap, path)
			}
			return nil
		}
	}
	return nil
}

// sendMsg sends the arg into the channel
func (s *pubSub) sendMsg(a arg) {
	if false == s.closed {
		s.publishArg <- a
	}
}

// publish runs in a Goroutine and listens on the channel publishArg. Every time
// a message is coming in, it calls all subscribers. We must run asynchronously
// because we don't know how long each subscriber needs.
func (s *pubSub) publish() {

	for {
		select {
		case <-s.stop:
			s.closeErr <- nil
			return
		case a, ok := <-s.publishArg:
			if !ok {
				// channel closed
				return
			}

			if len(s.subMap) == 0 {
				break
			}

			s.mu.RLock()
			var evict []int

			if subs, ok := s.subMap[a.pathLevel1()]; ok { // e.g.: system
				evict = append(evict, sendMessages(subs, a)...)
			}
			if subs, ok := s.subMap[a.pathLevel2()]; ok { // e.g.: system/smtp
				evict = append(evict, sendMessages(subs, a)...)
			}
			if subs, ok := s.subMap[a.pathLevelAll()]; ok { // e.g.: system/smtp/host
				evict = append(evict, sendMessages(subs, a)...)
			}
			s.mu.RUnlock()

			// remove all Subscribers which failed
			if len(evict) > 0 {
				for _, e := range evict {
					if err := s.Unsubscribe(e); err != nil && PkgLog.IsDebug() {
						PkgLog.Debug("config.pubSub.publish.evict.Unsubscribe.err", "err", err, "subscriptionID", e)
					}
				}
			}
		}
	}
}

func sendMessages(subs map[int]MessageReceiver, a arg) (evict []int) {
	for id, s := range subs {
		if err := sendMsgRecoverable(id, s, a); err != nil {
			if PkgLog.IsDebug() {
				PkgLog.Debug("config.pubSub.publish.sendMessages", "err", err, "id", id)
			}
			evict = append(evict, id) // mark Subscribers for removal which failed ...
		}
	}
	return
}

func sendMsgRecoverable(id int, sl MessageReceiver, a arg) (err error) {
	defer func() { // protect ... you'll never know
		if r := recover(); r != nil {
			if recErr, ok := r.(error); ok {
				PkgLog.Debug("config.pubSub.publish.recover.err", "err", recErr)
				err = recErr
			} else {
				PkgLog.Debug("config.pubSub.publish.recover.r", "recover", r)
				err = errgo.Newf("%#v", r)
			}
			// the overall trick here is, that defer will assign a new error to err
			// and therefore will overwrite the returned nil value!
		}
	}()
	err = sl.MessageConfig(a.pathLevelAll(), a.scope, a.scopeID)
	return
}

func newPubSub() *pubSub {
	return &pubSub{
		subMap:     make(map[string]map[int]MessageReceiver),
		publishArg: make(chan arg),
		stop:       make(chan struct{}),
		closeErr:   make(chan error),
	}
}
