// SPDX-License-Identifier: MIT

package validation

import (
	"testing"

	"github.com/issue9/assert"
)

func TestIn(t *testing.T) {
	a := assert.New(t)

	rule := In(1, 2, "3", struct{}{})
	a.False(rule.IsValid(3))
	a.False(rule.IsValid("1"))
	a.True(rule.IsValid(1))
	a.True(rule.IsValid(uint8(1)))

	rule = In(1, "2", &objectWithValidate{}, &objectWithValidate{Name: "name"})
	a.False(rule.IsValid(3))
	a.False(rule.IsValid("1"))
	a.True(rule.IsValid(&objectWithValidate{}))
	a.True(rule.IsValid(&objectWithValidate{Name: "name"}))
	a.False(rule.IsValid(&objectWithValidate{Name: "name", Age: 1}))
}

func TestNotIn(t *testing.T) {
	a := assert.New(t)

	rule := NotIn(1, 2, "3", struct{}{})
	a.True(rule.IsValid(3))
	a.True(rule.IsValid("1"))
	a.False(rule.IsValid(1))
	a.False(rule.IsValid(uint8(1)))

	rule = NotIn(1, "2", &objectWithoutValidate{}, &objectWithoutValidate{Name: "name"})
	a.True(rule.IsValid(3))
	a.True(rule.IsValid("1"))
	a.False(rule.IsValid(&objectWithoutValidate{}))
	a.False(rule.IsValid(&objectWithoutValidate{Name: "name"}))
	a.True(rule.IsValid(&objectWithoutValidate{Name: "name", Age: 1}))
}
