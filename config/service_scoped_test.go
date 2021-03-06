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

package config_test

import (
	"github.com/corestoreio/csfw/config"
	"github.com/corestoreio/csfw/store/scope"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestScopedServiceScope(t *testing.T) {
	tests := []struct {
		websiteID, groupID, storeID int64
		wantScope                   scope.Scope
		wantID                      int64
	}{
		{0, 0, 0, scope.DefaultID, 0},
		{1, 0, 0, scope.WebsiteID, 1},
		{1, 2, 0, scope.GroupID, 2},
		{1, 2, 3, scope.StoreID, 3},
		{0, 0, 3, scope.StoreID, 3},
		{0, 2, 0, scope.GroupID, 2},
	}
	for i, test := range tests {
		sg := config.NewMockGetter().NewScoped(test.websiteID, test.groupID, test.storeID)
		haveScope, haveID := sg.Scope()
		assert.Exactly(t, test.wantScope, haveScope, "Index %d", i)
		assert.Exactly(t, test.wantID, haveID, "Index %d", i)
	}
}

func TestScopedService(t *testing.T) {
	tests := []struct {
		desc                        string
		fqpath                      string
		path                        []string
		websiteID, groupID, storeID int64
		err                         error
	}{
		{
			"Default ScopedGetter should return default scope",
			scope.StrDefault.FQPath("0", "a/b/c"), []string{"a/b/c"}, 0, 0, 0, nil,
		},
		{
			"Website ID 1 ScopedGetter should fall back to default scope",
			scope.StrDefault.FQPath("0", "a/b/c"), []string{"a/b/c"}, 1, 0, 0, nil,
		},
		{
			"Website ID 10 + Group ID 12 ScopedGetter should fall back to website 10 scope",
			scope.StrWebsites.FQPath("10", "a/b/c"), []string{"a/b/c"}, 10, 12, 0, nil,
		},
		{
			"Website ID 10 + Group ID 12 + Store 22 ScopedGetter should fall back to website 10 scope",
			scope.StrWebsites.FQPath("10", "a/b/c"), []string{"a/b/c"}, 10, 12, 22, nil,
		},
		{
			"Website ID 10 + Group ID 12 + Store 22 ScopedGetter should return Store 22 scope",
			scope.StrStores.FQPath("22", "a/b/c"), []string{"a/b/c"}, 10, 12, 22, nil,
		},
		{
			"Website ID 10 + Group ID 12 + Store 42 ScopedGetter should return nothing",
			scope.StrStores.FQPath("22", "a/b/c"), []string{"a/b/c"}, 10, 12, 42, config.ErrKeyNotFound,
		},
		{
			"Path consists of only two elements which is incorrect",
			scope.StrDefault.FQPath("0", "a/b/c"), []string{"a", "b"}, 0, 0, 0, config.ErrPathEmpty,
		},
	}

	vals := []interface{}{"Gopher", true, float64(3.14159), int(2016), time.Now()}

	for vi, val := range vals {
		for _, test := range tests {

			cg := config.NewMockGetter(config.WithMockValues(config.MockPV{
				test.fqpath: val,
			}))

			sg := cg.NewScoped(test.websiteID, test.groupID, test.storeID)

			switch val.(type) {
			case string:
				s, err := sg.String(test.path...)
				testScopedService(t, s, test.desc, err, test.err)
			default:
				t.Fatalf("Unsupported type: %#v in index %d", val, vi)
			}
		}
	}
}

func testScopedService(t *testing.T, have, desc string, err, wantErr error) {
	if wantErr != nil {
		assert.Empty(t, have, desc)
		assert.EqualError(t, err, wantErr.Error(), desc)
		return
	}
	assert.NoError(t, err, desc)
	assert.Exactly(t, "Gopher", have, desc)
}

func BenchmarkScopedServiceStringDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
func BenchmarkScopedServiceStringWebsite(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
func BenchmarkScopedServiceStringGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
func BenchmarkScopedServiceStringStore(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
