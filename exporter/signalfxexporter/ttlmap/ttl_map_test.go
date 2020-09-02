// Copyright 2020, OpenTelemetry Authors
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

package ttlmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTTLMapData(t *testing.T) {
	m := newTTLMapData(10)
	require.Nil(t, m.get("foo"))
	m.put("bob", "xyz", 2)
	m.sweep(12)
	require.NotNil(t, m.get("bob"))
	m.sweep(13)
	require.Nil(t, m.get("bob"))
}

func TestTTLMap(t *testing.T) {
	ttlMap := New(5, 10)
	require.EqualValues(t, ttlMap.sweepInterval, 5)
	require.EqualValues(t, ttlMap.md.maxAge, 10)
	ttlMap.Put("foo", "bar")
	s := ttlMap.Get("foo").(string)
	require.Equal(t, "bar", s)
}
