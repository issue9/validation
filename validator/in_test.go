// SPDX-License-Identifier: MIT

package validator

import (
	"testing"

	"github.com/issue9/assert/v2"

	"github.com/issue9/validation"
)

type object struct {
	Name string
	Age  int
}

type objectWithoutFieldValidator struct {
	Name string
	Age  int
}

func (obj *object) ValidateFields(v *validation.Validation) {
	v.NewField(obj.Age, "age", Min(18).Message("不能小于 18"))
}

func TestIn(t *testing.T) {
	a := assert.New(t, false)

	rule := In(1, 2, "3", struct{}{})
	a.False(rule.IsValid(3))
	a.False(rule.IsValid("1"))
	a.True(rule.IsValid(1))
	a.True(rule.IsValid(uint8(1)))

	rule = In(1, "2", &object{}, &object{Name: "name"})
	a.False(rule.IsValid(3))
	a.False(rule.IsValid("1"))
	a.True(rule.IsValid(&object{}))
	a.True(rule.IsValid(&object{Name: "name"}))
	a.False(rule.IsValid(&object{Name: "name", Age: 1}))
}

func TestNotIn(t *testing.T) {
	a := assert.New(t, false)

	rule := NotIn(1, 2, "3", struct{}{})
	a.True(rule.IsValid(3))
	a.True(rule.IsValid("1"))
	a.False(rule.IsValid(1))
	a.False(rule.IsValid(uint8(1)))

	rule = NotIn(1, "2", &objectWithoutFieldValidator{}, &objectWithoutFieldValidator{Name: "name"})
	a.True(rule.IsValid(3))
	a.True(rule.IsValid("1"))
	a.False(rule.IsValid(&objectWithoutFieldValidator{}))
	a.False(rule.IsValid(&objectWithoutFieldValidator{Name: "name"}))
	a.True(rule.IsValid(&objectWithoutFieldValidator{Name: "name", Age: 1}))
}
