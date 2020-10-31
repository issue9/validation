// SPDX-License-Identifier: MIT

package validation

import (
	"testing"

	"github.com/issue9/assert"
)

func TestErrors(t *testing.T) {
	a := assert.New(t)

	errs := Errors{}

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
