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

import "fmt"

// ParseFunc is the signature for shard selector parsers.
type ParseFunc func(string) (Selector, error)

var registeredParser ParseFunc

// RegisterParser registers the shard selector parser implementation.
// This is called by k8s.io/apiserver to inject the CEL-based parser.
//
// This registration mechanism is an alpha-level internal API and is subject
// to change or removal in future releases. Do not depend on it.
func RegisterParser(p ParseFunc) {
	registeredParser = p
}

// Parse parses a shard selector string into a Selector.
// A parser must be registered via RegisterParser before calling Parse.
func Parse(s string) (Selector, error) {
	if registeredParser == nil {
		return nil, fmt.Errorf("no shard selector parser registered")
	}
	return registeredParser(s)
}
