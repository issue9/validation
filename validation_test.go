// SPDX-License-Identifier: MIT

package validation

import (
	"testing"

	"github.com/issue9/assert"
)

type (
	root struct {
		O1 *objectWithValidate
		O2 *objectWithoutValidate
	}

	objectWithValidate struct {
		Name string
		Age  int
	}

	objectWithoutValidate struct {
		Name string
		Age  int
	}
)

func (obj *objectWithValidate) Validate(errHandling ErrorHandling) Messages {
	return New(errHandling).
		NewField(obj.Age, ".age", Min("不能小于 18", 18)).
		Result()
}

func (root *root) Validate(errHandling ErrorHandling) Messages {
	return New(errHandling).
		NewField(root.O1, "o1", If(root.O2 == nil, Required("o1 required", true)).Rules()...).
		NewField(root.O2, "o2", If(root.O1 == nil, Required("o2 required", true)).Rules()...).
		Result()
}

func TestValidation_ErrorHandling(t *testing.T) {
	a := assert.New(t)

	v := New(ContinueAtError).
		NewField(-100, "f1", Min("-2", -2), Min("-3", -3)).
		NewField(100, "f2", Max("50", 50), Max("-4", -4))
	a.Equal(v.Result(), map[string][]string{
		"f1": {"-2", "-3"},
		"f2": {"50", "-4"},
	})

	v = New(ExitFieldAtError).
		NewField(-100, "f1", Min("-2", -2), Min("-3", -3)).
		NewField(100, "f2", Max("50", 50), Max("-4", -4))
	a.Equal(v.Result(), map[string][]string{
		"f1": {"-2"},
		"f2": {"50"},
	})

	v = New(ExitAtError).
		NewField(-100, "f1", Min("-2", -2), Min("-3", -3)).
		NewField(100, "f2", Max("50", 50), Max("-4", -4))
	a.Equal(v.Result(), map[string][]string{
		"f1": {"-2"},
	})
}

func TestValidation_NewObject(t *testing.T) {
	a := assert.New(t)

	obj := &objectWithValidate{}
	v := New(ContinueAtError).
		NewField(obj, "obj")
	a.Equal(v.Result(), map[string][]string{
		"obj.age": {"不能小于 18"},
	})

	//
	r := root{}
	errs := r.Validate(ContinueAtError)
	a.Equal(errs, map[string][]string{
		"o1": {"o1 required"},
		"o2": {"o2 required"},
	})

	r = root{O1: &objectWithValidate{}}
	errs = r.Validate(ContinueAtError)
	a.Equal(errs, map[string][]string{
		"o1.age": {"不能小于 18"},
	})
}
