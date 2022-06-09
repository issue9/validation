// SPDX-License-Identifier: MIT

package validator

import (
	"testing"

	"github.com/issue9/assert/v2"
)

type object struct {
	Name string
	Age  int
}

func TestIn(t *testing.T) {
	a := assert.New(t, false)

	rule := In(1, 2)
	a.False(rule.IsValid(3))
	a.False(rule.IsValid("1"))
	a.True(rule.IsValid(1))
	a.False(rule.IsValid(uint8(1)))

	rule = In(object{}, object{Name: "name"})
	a.False(rule.IsValid(3))
	a.False(rule.IsValid("1"))
	a.True(rule.IsValid(object{}))
	a.True(rule.IsValid(object{Name: "name"}))
	a.False(rule.IsValid(object{Name: "name", Age: 1}))

	rule = In(&object{}, &object{Name: "name"})
	a.False(rule.IsValid(&object{}))
	a.False(rule.IsValid(&object{Name: "name"}))
	a.False(rule.IsValid(&object{Name: "name", Age: 1}))
}

func TestNotIn(t *testing.T) {
	a := assert.New(t, false)

	rule := NotIn(1, 2)
	a.True(rule.IsValid(3))
	a.True(rule.IsValid("1"))
	a.False(rule.IsValid(1))
	a.True(rule.IsValid(uint8(1)))

	rule = NotIn(object{}, object{Name: "name"})
	a.True(rule.IsValid(3))
	a.True(rule.IsValid("1"))
	a.False(rule.IsValid(object{}))
	a.False(rule.IsValid(object{Name: "name"}))
	a.True(rule.IsValid(object{Name: "name", Age: 1}))

	rule = NotIn(&object{}, &object{Name: "name"})
	a.True(rule.IsValid(&object{}))
	a.True(rule.IsValid(&object{Name: "name"}))
	a.True(rule.IsValid(&object{Name: "name", Age: 1}))
}
