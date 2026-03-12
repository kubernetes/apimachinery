/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sharding

import (
	"testing"
)

func TestParseWithoutRegisteredParser(t *testing.T) {
	old := registeredParser
	defer func() { registeredParser = old }()

	registeredParser = nil
	_, err := Parse("anything")
	if err == nil {
		t.Error("Parse() without registered parser should return error")
	}
}

func TestRegisterParserAndParse(t *testing.T) {
	old := registeredParser
	defer func() { registeredParser = old }()

	called := false
	RegisterParser(func(s string) (Selector, error) {
		called = true
		return NewSelector(ShardRangeRequirement{
			Key:   "object.metadata.uid",
			Start: "0x0",
			End:   "0x8",
		}), nil
	})

	sel, err := Parse("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !called {
		t.Error("registered parser was not called")
	}
	reqs := sel.Requirements()
	if len(reqs) != 1 || reqs[0].Key != "object.metadata.uid" {
		t.Errorf("unexpected requirements: %+v", reqs)
	}
}
