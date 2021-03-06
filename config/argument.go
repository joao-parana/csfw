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
	"fmt"
	"io"
	"io/ioutil"

	"github.com/corestoreio/csfw/store/scope"
	"github.com/corestoreio/csfw/util"
)

const hierarchyLevel int = 3 // a/b/c

// ErrPathEmpty when you provide an empty path in the function Path()
var ErrPathEmpty = errors.New("Path cannot be empty")

// ArgFunc Argument function to be used as variadic argument in ScopeKey() and ScopeKeyValue()
type ArgFunc func(*arg)

// ScopeDefault wrapper helper function. See Scope(). Mainly used to show humans
// than a config value can only be set for a global scope.
func ScopeDefault() ArgFunc { return Scope(scope.DefaultID, 0) }

// ScopeWebsite wrapper helper function. See Scope()
func ScopeWebsite(id int64) ArgFunc { return Scope(scope.WebsiteID, id) }

// ScopeGroup wrapper helper function. See Scope()
func ScopeGroup(id int64) ArgFunc { return Scope(scope.GroupID, id) }

// ScopeStore wrapper helper function. See Scope()
func ScopeStore(id int64) ArgFunc { return Scope(scope.StoreID, id) }

// Scope sets the scope using the scope.Group and a ID.
// The ID can contain an integer from a website or a store. Make sure
// the correct scope.Scope has also been set. If the ID is smaller
// than zero the scope will fallback to default scope.
func Scope(s scope.Scope, id int64) ArgFunc {
	if s != scope.DefaultID && id < 1 {
		id = 0
		s = scope.DefaultID
	}
	return func(a *arg) { a.scope = s; a.scopeID = id }
}

// Path option function to specify the configuration path. If one argument has been
// provided then it must be a full valid path. If more than one argument has been provided
// then the arguments will be joined together. Panics if nil arguments will be provided.
func Path(paths ...string) ArgFunc {
	// TODO(cs) validation of the path see typeConfigPath in app/code/Magento/Config/etc/system_file.xsd

	if false == isValidPath(paths...) {
		return func(a *arg) {
			a.lastErrors = append(a.lastErrors, ErrPathEmpty)
		}
	}

	var paSlice []string
	if len(paths) >= hierarchyLevel {
		paSlice = paths
	} else {
		paSlice = scope.PathSplit(paths[0])
		if len(paSlice) < hierarchyLevel {
			return func(a *arg) {
				a.lastErrors = append(a.lastErrors, fmt.Errorf("Incorrect number of paths elements: want %d, have %d, Path: %v", hierarchyLevel, len(paSlice), paths))
			}
		}
	}
	return func(a *arg) {
		a.pathSlice = paSlice
	}
}

// Value sets the value for a scope key.
func Value(v interface{}) ArgFunc { return func(a *arg) { a.v = v } }

// ValueReader sets the value for a scope key using the io.Reader interface.
// If asserting to a io.Closer is successful then Close() will be called.
func ValueReader(r io.Reader) ArgFunc {
	data, err := ioutil.ReadAll(r)
	if c, ok := r.(io.Closer); ok && c != nil {
		if err := c.Close(); err != nil {
			return func(a *arg) {
				a.lastErrors = append(a.lastErrors, fmt.Errorf("ValueReader.Close error %s", err))
			}
		}
	}
	if err != nil {
		return func(a *arg) {
			a.lastErrors = append(a.lastErrors, fmt.Errorf("ValueReader error %s", err))
		}
	}
	return func(a *arg) {
		a.v = data
	}
}

// isValidPath checks for valid config path. Either full path like general/country/allow
// or at least 3 path parts.
func isValidPath(paths ...string) bool {
	return (len(paths) == 1 && paths[0] != "") ||
		(len(paths) >= hierarchyLevel && paths[0] != "" && paths[1] != "" && paths[2] != "")
}

// arg responsible for the correct scope key e.g.: stores/2/system/currency/installed => scope/scope_id/path
// which is used by the underlying configuration Service to fetch or store a value
type arg struct {
	pathSlice  []string // pa is the three level path e.g. a/b/c split by slash
	scope      scope.Scope
	scopeID    int64       // scope ID
	v          interface{} // value use for saving
	lastErrors []error
}

// newArg creates an argument container which requires different options depending on the use case.
func newArg(opts ...ArgFunc) (arg, error) {
	var a = arg{}
	for _, opt := range opts {
		if opt != nil {
			opt(&a)
		}
	}
	if len(a.lastErrors) > 0 {
		return arg{}, a
	}
	return a, nil
}

// mustNewArg panics on error. useful for initialization process
func mustNewArg(opts ...ArgFunc) arg {
	a, err := newArg(opts...)
	if err != nil {
		panic(err)
	}
	return a
}

func (a arg) isValidPath() bool    { return isValidPath(a.pathSlice...) }
func (a arg) isDefault() bool      { return a.scope == scope.DefaultID || a.scope == scope.AbsentID }
func (a arg) pathLevel1() string   { return a.pathSlice[0] }
func (a arg) pathLevel2() string   { return scope.PathJoin(a.pathSlice[:2]...) }
func (a arg) pathLevelAll() string { return scope.PathJoin(a.pathSlice...) }

func (a arg) scopePath() string {
	// first part of the path is called scope in Magento and in CoreStore ScopeRange
	// e.g.: stores/2/system/currency/installed => scope/scope_id/path
	// e.g.: websites/1/system/currency/installed => scope/scope_id/path
	if false == a.isValidPath() {
		return ""
	}
	return scope.FromScope(a.scope).FQPathInt64(a.scopeID, a.pathSlice...)
}

// scopePathDefault returns a path prefixed by default StrScope
// e.g.: default/0/system/currency/installed
// 		 scope/scope_id/path...
//func (a arg) scopePathDefault() string { return scope.StrDefault.FQPath("0", a.pathSlice...) }

var _ error = (*arg)(nil)

func (a arg) Error() string {
	return util.Errors(a.lastErrors...)
}
