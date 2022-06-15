// SPDX-License-Identifier: MIT

package validator

import (
	"regexp"
	"testing"

	"github.com/issue9/assert/v2"
)

func TestMatch(t *testing.T) {
	a := assert.New(t, false)

	r := Match(regexp.MustCompile("[a-z]+"))
	a.True(r.IsValid("abc"))
	a.True(r.IsValid([]byte("def")))
	a.False(r.IsValid([]rune("123")))
	a.False(r.IsValid(123)) // 无法验证
}

func TestRequired(t *testing.T) {
	a := assert.New(t, false)
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
