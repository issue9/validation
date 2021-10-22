// SPDX-License-Identifier: MIT

package validation_test

import (
	"testing"

	"github.com/issue9/assert"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/issue9/validation"
	"github.com/issue9/validation/validator"
)

var _ validation.Validator = validator.Max(100)

func TestRule_AsSlice(t *testing.T) {
	a := assert.New(t)

	// 将数组当普通元素处理
	v := validation.New(validation.ContinueAtError, message.NewPrinter(language.SimplifiedChinese), "/").
		NewField([]int{1, 2, 6}, "slice", validator.Min(5).Message("min-5"))
	a.Equal(v.Messages(), validation.Messages{
		"slice": []string{"min-5"},
	})

	// 普通元素指定为 asSlice
	v = validation.New(validation.ContinueAtError, message.NewPrinter(language.SimplifiedChinese), "/").
		NewField("123456", "slice", validator.Min(5).Message("min-5").AsSlice())
	a.Equal(v.Messages(), validation.Messages{
		"slice": []string{"min-5"},
	})

	// ContinueAtError
	v = validation.New(validation.ContinueAtError, message.NewPrinter(language.SimplifiedChinese), "/").
		NewField([]int{1, 2, 6}, "slice", validator.Min(5).Message("min-5").AsSlice())
	a.Equal(v.Messages(), validation.Messages{
		"slice[0]": []string{"min-5"},
		"slice[1]": []string{"min-5"},
	})

	// ExitAtError
	v = validation.New(validation.ExitAtError, message.NewPrinter(language.SimplifiedChinese), "/").
		NewField([]int{1, 2, 6}, "slice", validator.Min(5).Message("min-5").AsSlice())
	a.Equal(v.Messages(), validation.Messages{
		"slice[0]": []string{"min-5"},
	})
}
