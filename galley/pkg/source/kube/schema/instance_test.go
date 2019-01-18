// Copyright 2018 Istio Authors
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

package schema

import (
	"reflect"
	"testing"
)

func TestSchemaBuilder(t *testing.T) {
	spec := ResourceSpec{
		Kind:     "kind",
		Version:  "version",
		ListKind: "listkind",
		Plural:   "plural",
		Singular: "singular",
		Group:    "groupd",
	}

	b := NewBuilder()
	b.Add(spec)
	s := b.Build()

	actual := s.All()

	expected := []ResourceSpec{spec}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Mismatch:\nGot:\n%v\nWanted:\n%v\n", actual, expected)
	}
}