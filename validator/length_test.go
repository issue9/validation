// SPDX-License-Identifier: MIT

package validator

import (
	"testing"

	"github.com/issue9/assert/v2"
)

func TestLength(t *testing.T) {
	a := assert.New(t, false)

	a.Panic(func() {
		Length(500, 50)
	})

	l := Length(5, 7)
	a.False(l.IsValid("123"))
	a.False(l.IsValid([]byte("123")))
	a.True(l.IsValid([]rune("12345")))
	a.False(l.IsValid(&struct{}{}))

	// 不限制长度
	l = Length(-1, -1)
	a.True(l.IsValid("12345678910"))
	a.True(l.IsValid([]rune("")))

	l = MinLength(6)
	a.True(l.IsValid("123456"))
	a.True(l.IsValid("12345678910"))
	a.False(l.IsValid("12345"))

	l = MaxLength(6)
	a.True(l.IsValid("123456"))
	a.False(l.IsValid("12345678910"))
	a.True(l.IsValid("12345"))
}
