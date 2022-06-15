// SPDX-License-Identifier: MIT

package validation_test

import (
	"testing"

	"github.com/issue9/assert/v2"
	"github.com/issue9/localeutil"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/issue9/validation"
	"github.com/issue9/validation/validator"
)

var _ validation.Validator = validator.Max(100)

func TestRule_AsSlice(t *testing.T) {
	a := assert.New(t, false)
	p := message.NewPrinter(language.SimplifiedChinese)

	// 将数组当普通元素处理
	v := validation.New(validation.ContinueAtError, 10).
		NewField([]int{1, 2, 6}, "slice", validator.Min(5).Message("min-5"))
	a.Equal(v.Messages(), validation.Messages{
		"slice": []localeutil.LocaleStringer{localeutil.Phrase("min-5")},
	})

	// 普通元素指定为 asSlice
	v = validation.New(validation.ContinueAtError, 10).
		NewField("123456", "slice", validator.Min(5).Message("min-5").AsSlice())
	a.Equal(v.LocaleMessages(p), validation.LocaleMessages{
		"slice": []string{"min-5"},
	})

	// ContinueAtError
	v = validation.New(validation.ContinueAtError, 10).
		NewField([]int{1, 2, 6}, "slice", validator.Min(5).Message("min-5").AsSlice())
	a.Equal(v.LocaleMessages(p), validation.LocaleMessages{
		"slice[0]": []string{"min-5"},
		"slice[1]": []string{"min-5"},
	})

	// ExitAtError
	v = validation.New(validation.ExitAtError, 10).
		NewField([]int{1, 2, 6}, "slice", validator.Min(5).Message("min-5").AsSlice())
	a.Equal(v.LocaleMessages(p), validation.LocaleMessages{
		"slice[0]": []string{"min-5"},
	})
}

func TestIfExpr(t *testing.T) {
	a := assert.New(t, false)

	count := 5
	rules := validation.If(count == 0, validator.Min(5).Message("min-5")).Else(validator.Max(100).Message("max-100"), validator.Min(50).Message("min-50")).Rules()
	a.Equal(2, len(rules))

	// 第二次 Else 清空了之前的规则
	count = 5
	rules = validation.If(count == 0, validator.Min(5).Message("min-5")).Else(validator.Max(100).Message("max-100"), validator.Min(50).Message("min-50")).Else().Rules()
	a.Equal(0, len(rules))

	count = 0
	rules = validation.If(count == 0, validator.Min(5).Message("min-5")).Else(validator.Max(100).Message("max-100"), validator.Min(50).Message("min-50")).Rules()
	a.Equal(1, len(rules))
}
