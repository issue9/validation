// SPDX-License-Identifier: MIT

package validation

import (
	"testing"

	"github.com/issue9/assert"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type (
	root1 struct {
		Root *root2
		F1   int
	}

	root2 struct {
		O1 *object
		O2 *objectWithoutFieldValidator
	}

	object struct {
		Name string
		Age  int
	}

	objectWithoutFieldValidator struct {
		Name string
		Age  int
	}
)

func (obj *object) ValidateFields(v *Validation) {
	v.NewField(obj.Age, "age", Min(18).Message("不能小于 18"))
}

func (root *root2) ValidateFields(v *Validation) {
	v.NewField(root.O1, "o1", If(root.O2 == nil, Required(true).Message("o1 required")).Rules()...).
		NewField(root.O2, "o2", If(root.O1 == nil, Required(true).Message("o2 required")).Rules()...)
}

func (root *root1) ValidateFields(v *Validation) {
	v.NewField(root.Root, "root", Required(false).Message("root required")).
		NewField(root.F1, "f1", Min(5).Message("min-5"))
}

func TestValidation_ErrorHandling(t *testing.T) {
	a := assert.New(t)

	v := New(ContinueAtError, message.NewPrinter(language.Chinese), "/").
		NewField(-100, "f1", Min(-2).Message("-2"), Min(-3).Message("-3")).
		NewField(100, "f2", Max(50).Message("50"), Max(-4).Message("-4"))
	a.Equal(v.Messages(), Messages{
		"f1": {"-2", "-3"},
		"f2": {"50", "-4"},
	})

	v = New(ExitFieldAtError, message.NewPrinter(language.Chinese), "/").
		NewField(-100, "f1", Min(-2).Message("-2"), Min(-3).Message("-3")).
		NewField(100, "f2", Max(50).Message("50"), Max(-4).Message("-4"))
	a.Equal(v.Messages(), Messages{
		"f1": {"-2"},
		"f2": {"50"},
	})

	v = New(ExitAtError, message.NewPrinter(language.Chinese), "/").
		NewField(-100, "f1", Min(-2).Message("-2"), Min(-3).Message("-3")).
		NewField(100, "f2", Max(50).Message("50"), Max(-4).Message("-4"))
	a.Equal(v.Messages(), Messages{
		"f1": {"-2"},
	})
}

func TestValidation_NewField(t *testing.T) {
	a := assert.New(t)

	obj := &object{}
	v := New(ContinueAtError, message.NewPrinter(language.Chinese), "/").
		NewField(obj, "obj")
	a.Equal(v.Messages(), Messages{
		"obj/age": {"不能小于 18"},
	})

	// object
	r := root2{}
	v = New(ContinueAtError, message.NewPrinter(language.Chinese), "/")
	r.ValidateFields(v)
	a.Equal(v.Messages(), Messages{
		"o1": {"o1 required"},
		"o2": {"o2 required"},
	})

	r = root2{O1: &object{}}
	v = New(ContinueAtError, message.NewPrinter(language.Chinese), ".")
	r.ValidateFields(v)
	a.Equal(v.Messages(), Messages{
		"o1.age": {"不能小于 18"},
	})

	v = New(ContinueAtError, message.NewPrinter(language.Chinese), "/")
	rv := root1{Root: &root2{O1: &object{}}}
	rv.ValidateFields(v)
	a.Equal(v.Messages(), Messages{
		"root/o1/age": {"不能小于 18"},
		"f1":          {"min-5"},
	})

	// slice
	v = New(ContinueAtError, message.NewPrinter(language.SimplifiedChinese), "/")
	messages := v.NewField([]int{1, 2, 6}, "slice", Min(5).Message("min-5")).Messages()
	a.Equal(messages, Messages{
		"slice": []string{"min-5"},
	})

	v = New(ContinueAtError, message.NewPrinter(language.SimplifiedChinese), "/")
	messages = v.NewField([]int{1, 2, 6}, "slice", Min(5).Message("min-5").AsSlice()).Messages()
	a.Equal(messages, Messages{
		"slice[0]": []string{"min-5"},
		"slice[1]": []string{"min-5"},
	})
}

func TestValidation_Locale(t *testing.T) {
	a := assert.New(t)
	builder := catalog.NewBuilder()
	builder.SetString(language.SimplifiedChinese, "lang", "chn")
	builder.SetString(language.TraditionalChinese, "lang", "cht")

	v := New(ContinueAtError, message.NewPrinter(language.SimplifiedChinese, message.Catalog(builder)), "/").
		NewField(5, "obj", Max(4).Message("lang"))
	a.Equal(v.Messages(), Messages{
		"obj": {"chn"},
	})

	v = New(ContinueAtError, message.NewPrinter(language.TraditionalChinese, message.Catalog(builder)), "/").
		NewField(5, "obj", Max(4).Message("lang"))
	a.Equal(v.Messages(), Messages{
		"obj": {"cht"},
	})
}
