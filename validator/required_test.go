// SPDX-License-Identifier: MIT

package validator

import (
	"testing"

	"github.com/issue9/assert"
)

func TestRequired(t *testing.T) {
	a := assert.New(t)
	val := 5

	r := Required(false)
	a.False(r.IsValid(0))
	a.False(r.IsValid(nil))
	a.False(r.IsValid(""))
	a.False(r.IsValid([]string{}))
	a.True(r.IsValid([]string{""}))
	a.True(r.IsValid(&val))

	r = Required(true)
	a.False(r.IsValid(0))
	a.True(r.IsValid(nil))
	a.False(r.IsValid(""))
	a.False(r.IsValid([]string{}))
	a.True(r.IsValid([]string{""}))
	a.True(r.IsValid(&val))
}
