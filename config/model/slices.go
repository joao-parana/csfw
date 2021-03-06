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

package model

import (
	"strconv"
	"strings"

	"github.com/corestoreio/csfw/config"
	"github.com/corestoreio/csfw/store/scope"
	"github.com/corestoreio/csfw/util/bufferpool"
	"github.com/juju/errgo"
)

// CSVSeparator separates CSV values
const CSVSeparator = ","

// StringCSV represents a path in config.Getter which will be saved as a
// CSV string and returned as a string slice. Separator is a comma.
type StringCSV struct{ basePath }

// NewStringCSV creates a new CSV string type. Acts as a multiselect.
func NewStringCSV(path string, opts ...Option) StringCSV {
	return StringCSV{basePath: NewPath(path, opts...)}
}

// Get returns a string slice
func (p StringCSV) Get(sg config.ScopedGetter) []string {
	s, err := p.lookupString(sg)
	if err != nil && PkgLog.IsDebug() {
		PkgLog.Debug("model.StringCSV.Get.lookupString", "err", err, "path", p.string)
	}
	if s == "" {
		return nil
	}
	// validate ?
	return strings.Split(s, CSVSeparator)
}

// Write writes a slice with its scope and ID to the writer
func (p StringCSV) Write(w config.Writer, sl []string, s scope.Scope, id int64) error {
	for _, v := range sl {
		if err := p.validateString(v); err != nil {
			return err
		}
	}
	return p.basePath.Write(w, strings.Join(sl, CSVSeparator), s, id)
}

// IntCSV represents a path in config.Getter which will be saved as a
// CSV string and returned as an int64 slice. Separator is a comma.
type IntCSV struct{ basePath }

// NewIntCSV creates a new int CSV type. Acts as a multiselect.
func NewIntCSV(path string, opts ...Option) IntCSV {
	return IntCSV{basePath: NewPath(path, opts...)}
}

func (p IntCSV) Get(sg config.ScopedGetter) []int {
	s, err := p.lookupString(sg)
	if err != nil && PkgLog.IsDebug() {
		PkgLog.Debug("model.IntCSV.Get.lookupString", "err", err, "path", p.string)
	}
	if s == "" {
		return nil
	}

	csv := strings.Split(s, CSVSeparator)
	ret := make([]int, 0, len(csv))
	for i, line := range csv {
		var err error

		v, err := strconv.Atoi(line)
		if err != nil {
			if PkgLog.IsDebug() {
				PkgLog.Debug("model.IntCSV.Get.strconv.ParseInt", "err", err, "position", i, "line", line)
			}
			continue
		}

		if err := p.validateInt(v); err != nil {
			if PkgLog.IsDebug() {
				PkgLog.Debug("model.IntCSV.Get.validateInt", "err", err, "position", i, "line", line)
			}
			continue
		}

		ret = append(ret, v)

	}
	return ret
}

// Write writes int values as a CSV string
func (p IntCSV) Write(w config.Writer, sl []int, s scope.Scope, id int64) error {

	val := bufferpool.Get()
	defer bufferpool.Put(val)
	for i, v := range sl {

		if err := p.validateInt(v); err != nil {
			return err
		}

		if _, err := val.WriteString(strconv.Itoa(v)); err != nil {
			return errgo.Mask(err)
		}
		if i < len(sl)-1 {
			if _, err := val.WriteString(CSVSeparator); err != nil {
				return errgo.Mask(err)
			}
		}
	}
	return p.basePath.Write(w, val.String(), s, id)
}
