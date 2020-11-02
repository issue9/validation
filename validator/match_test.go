// SPDX-License-Identifier: MIT

package validator

import (
	"regexp"
	"testing"

	"github.com/issue9/assert"
)

func TestMatch(t *testing.T) {
	a := assert.New(t)

	r := Match(regexp.MustCompile("[a-z]+"))
	a.True(r.IsValid("abc"))
	a.True(r.IsValid([]byte("def")))
	a.False(r.IsValid([]rune("123")))
	a.False(r.IsValid(123)) // 无法验证
}
