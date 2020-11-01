// SPDX-License-Identifier: MIT

package validation

import (
	"testing"

	"github.com/issue9/assert"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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

func (obj *objectWithValidate) ValidateFields(errHandling ErrorHandling, p *message.Printer) Messages {
	return New(errHandling, p).
		NewField(obj.Age, ".age", Min(18).Message("不能小于 18")).
		Messages()
}

func (root *root) ValidateFields(errHandling ErrorHandling, p *message.Printer) Messages {
	return New(errHandling, p).
		NewField(root.O1, "o1", If(root.O2 == nil, Required(true).Message("o1 required")).Rules()...).
		NewField(root.O2, "o2", If(root.O1 == nil, Required(true).Message("o2 required")).Rules()...).
		Messages()
}

func TestValidation_ErrorHandling(t *testing.T) {
	a := assert.New(t)

	v := New(ContinueAtError, message.NewPrinter(language.Chinese)).
		NewField(-100, "f1", Min(-2).Message("-2"), Min(-3).Message("-3")).
		NewField(100, "f2", Max(50).Message("50"), Max(-4).Message("-4"))
	a.Equal(v.Messages(), Messages{
		"f1": {"-2", "-3"},
		"f2": {"50", "-4"},
	})

	v = New(ExitFieldAtError, message.NewPrinter(language.Chinese)).
		NewField(-100, "f1", Min(-2).Message("-2"), Min(-3).Message("-3")).
		NewField(100, "f2", Max(50).Message("50"), Max(-4).Message("-4"))
	a.Equal(v.Messages(), Messages{
		"f1": {"-2"},
		"f2": {"50"},
	})

	v = New(ExitAtError, message.NewPrinter(language.Chinese)).
		NewField(-100, "f1", Min(-2).Message("-2"), Min(-3).Message("-3")).
		NewField(100, "f2", Max(50).Message("50"), Max(-4).Message("-4"))
	a.Equal(v.Messages(), Messages{
		"f1": {"-2"},
	})
}

func TestValidation_NewObject(t *testing.T) {
	a := assert.New(t)

	obj := &objectWithValidate{}
	v := New(ContinueAtError, message.NewPrinter(language.Chinese)).
		NewField(obj, "obj")
	a.Equal(v.Messages(), Messages{
		"obj.age": {"不能小于 18"},
	})

	//
	r := root{}
	errs := r.ValidateFields(ContinueAtError, message.NewPrinter(language.Chinese))
	a.Equal(errs, Messages{
		"o1": {"o1 required"},
		"o2": {"o2 required"},
	})

	r = root{O1: &objectWithValidate{}}
	errs = r.ValidateFields(ContinueAtError, message.NewPrinter(language.Chinese))
	a.Equal(errs, Messages{
		"o1.age": {"不能小于 18"},
	})
}
