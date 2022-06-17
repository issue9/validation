// SPDX-License-Identifier: MIT

package validation

import (
	"testing"

	"github.com/issue9/assert/v2"
	"github.com/issue9/localeutil"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"

	"github.com/issue9/validation/validator"
)

type (
	root1 struct {
		Root *root2
		F1   int
	}

	root2 struct {
		O1 *object
		O2 *object
	}

	object struct {
		Name string
		Age  int
	}
)

func TestValidation_ErrorHandling(t *testing.T) {
	a := assert.New(t, false)
	p := message.NewPrinter(language.Chinese)

	min_2 := NewRule(validator.Min(-2), "-2")
	min_3 := NewRule(validator.Min(-3), "-3")
	max50 := NewRule(validator.Max(50), "50")
	max_4 := NewRule(validator.Max(-4), "-4")

	v := New(ContinueAtError, 0).
		NewField(-100, "f1", min_2, min_3).
		NewField(100, "f2", max50, max_4)
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"f1": {"-2", "-3"},
		"f2": {"50", "-4"},
	})

	v = New(ExitFieldAtError, 1).
		NewField(-100, "f1", min_2, min_3).
		NewField(100, "f2", max50, max_4)
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"f1": {"-2"},
		"f2": {"50"},
	})

	v = New(ExitAtError, 1).
		NewField(-100, "f1", min_2, min_3).
		NewField(100, "f2", max50, max_4)
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"f1": {"-2"},
	})
}

func TestValidation_NewField(t *testing.T) {
	a := assert.New(t, false)
	p := message.NewPrinter(language.Chinese)

	min18 := NewRule(validator.Min(18), "不能小于 18")
	min5 := NewRule(validator.Min(5), "min-5")

	obj := &object{}
	v := New(ContinueAtError, 1).
		NewField(obj.Age, "obj/age", min18)
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"obj/age": {"不能小于 18"},
	})

	// object
	r := root2{}
	v = New(ContinueAtError, 10)
	v.NewField(r.O1, "o1", NewRule(validator.Required(false), "o1 required")).
		NewField(r.O2, "o2", NewRule(validator.Required(false), "o2 required"))
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"o1": {"o1 required"},
		"o2": {"o2 required"},
	})

	r = root2{O1: &object{}}
	v = New(ContinueAtError, 10)
	v.NewField(r.O1.Age, "o1.age", min18)
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"o1.age": {"不能小于 18"},
	})

	v = New(ContinueAtError, 10)
	rv := root1{Root: &root2{O1: &object{}}}
	v.NewField(rv.Root.O1.Age, "root/o1/age", min18).
		NewField(rv.F1, "f1", min5)
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"root/o1/age": {"不能小于 18"},
		"f1":          {"min-5"},
	})
}

func TestValidation_NewSliceField(t *testing.T) {
	a := assert.New(t, false)
	p := message.NewPrinter(language.SimplifiedChinese)

	min5 := NewRule(validator.Min(5), "min-5")

	// 将数组当普通元素处理
	v := New(ContinueAtError, 10).
		NewField([]int{1, 2, 6}, "slice", min5)
	a.Equal(v.Messages(), Messages{
		"slice": []localeutil.LocaleStringer{localeutil.Phrase("min-5")},
	})

	// 普通元素指定为 slice
	v = New(ContinueAtError, 10).
		NewSliceField(123456, "slice", min5)
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"slice": []string{"min-5"},
	})
	v = New(ExitAtError, 10).
		NewSliceField(123456, "slice", min5)
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"slice": []string{"min-5"},
	})

	// ContinueAtError
	v = New(ContinueAtError, 10).
		NewSliceField([]int{1, 2, 6}, "slice", min5)
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"slice[0]": []string{"min-5"},
		"slice[1]": []string{"min-5"},
	})

	// ExitAtError
	v = New(ExitAtError, 10).
		NewSliceField([]int{1, 2, 6}, "slice", min5)
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"slice[0]": []string{"min-5"},
	})
}

func TestValidation_NewMapField(t *testing.T) {
	a := assert.New(t, false)
	p := message.NewPrinter(language.SimplifiedChinese)

	min5 := NewRule(validator.Min(5), "min-5")

	// 将数组当普通元素处理
	v := New(ContinueAtError, 10).
		NewField([]int{1, 2, 6}, "slice", min5)
	a.Equal(v.Messages(), Messages{
		"slice": []localeutil.LocaleStringer{localeutil.Phrase("min-5")},
	})

	// 普通元素指定为 slice
	v = New(ContinueAtError, 10).
		NewMapField(123456, "slice", min5)
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"slice": []string{"min-5"},
	})
	v = New(ExitAtError, 10).
		NewMapField(123456, "slice", min5)
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"slice": []string{"min-5"},
	})

	// ContinueAtError
	v = New(ContinueAtError, 10).
		NewMapField(map[string]int{"0": 1, "2": 2, "6": 6}, "map", min5)
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"map[0]": []string{"min-5"},
		"map[2]": []string{"min-5"},
	})

	// ExitAtError
	v = New(ExitAtError, 10).
		NewMapField(map[string]int{"0": 1, "2": 2, "6": 6}, "map", min5)
	a.Length(v.LocaleMessages(p), 1) // map 顺序未定
}

func TestValidation_When(t *testing.T) {
	a := assert.New(t, false)
	p := message.NewPrinter(language.Chinese)

	min18 := NewRule(validator.Min(18), "不能小于 18")
	notEmpty := NewRule(validator.Required(true), "不能为空")

	obj := &object{}
	v := New(ContinueAtError, 1).
		NewField(obj, "obj/age", min18).
		When(obj.Age > 0, func(v *Validation) {
			v.NewField(obj.Name, "obj/name", notEmpty)
		})
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"obj/age": {"不能小于 18"},
	})

	obj = &object{Age: 15}
	v = New(ContinueAtError, 1).
		NewField(obj, "obj/age", min18).
		When(obj.Age > 0, func(v *Validation) {
			v.NewField(obj.Name, "obj/name", notEmpty)
		})
	a.Equal(v.LocaleMessages(p), LocaleMessages{
		"obj/age":  {"不能小于 18"},
		"obj/name": {"不能为空"},
	})
}

func TestValidation_Locale(t *testing.T) {
	a := assert.New(t, false)
	builder := catalog.NewBuilder()
	a.NotError(builder.SetString(language.SimplifiedChinese, "lang", "chn"))
	a.NotError(builder.SetString(language.TraditionalChinese, "lang", "cht"))

	max4 := NewRule(validator.Max(4), "lang")

	v := New(ContinueAtError, 10).
		NewField(5, "obj", max4)
	a.Equal(v.LocaleMessages(message.NewPrinter(language.SimplifiedChinese, message.Catalog(builder))), LocaleMessages{
		"obj": {"chn"},
	})

	v = New(ContinueAtError, 10).
		NewField(5, "obj", max4)
	a.Equal(v.LocaleMessages(message.NewPrinter(language.TraditionalChinese, message.Catalog(builder))), LocaleMessages{
		"obj": {"cht"},
	})
}
