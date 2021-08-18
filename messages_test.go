// SPDX-License-Identifier: MIT

package validation

import (
	"testing"

	"github.com/issue9/assert"
)

func TestMessages(t *testing.T) {
	a := assert.New(t)

	errs := Messages{}

	a.Panic(func() {
		errs.Add("key")
	})

	a.Panic(func() {
		errs.Set("key")
	})

	errs.Add("key1", "v1", "v2")
	a.Equal(errs, map[string][]string{"key1": {"v1", "v2"}})

	errs.Add("key1", "v1", "v3")
	a.Equal(errs, map[string][]string{"key1": {"v1", "v2", "v1", "v3"}})

	errs.Set("key1", "v1")
	a.Equal(errs, map[string][]string{"key1": {"v1"}})
}

func TestMessages_Merge(t *testing.T) {
	a := assert.New(t)

	m1 := Messages{}
	m1.Add("key1", "v1", "v2")

	m2 := Messages{}
	m2.Add("key1", "v2", "v3")
	m2.Add("key2", "v1")

	m1.Merge(m2)
	a.Equal(m1["key1"], []string{"v1", "v2", "v2", "v3"})
	a.Equal(m1["key2"], []string{"v1"})
}
