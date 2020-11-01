// SPDX-License-Identifier: MIT

package validation

import (
	"testing"

	"github.com/issue9/assert"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var _ Validator = Max(100)

func TestRule_valid(t *testing.T) {
	a := assert.New(t)

	// 将数组当普通元素处理
	v := New(ContinueAtError, message.NewPrinter(language.SimplifiedChinese))
	Min(5).Message("min-5").valid(v, "slice", []int{1, 2, 6})
	a.Equal(v.Messages(), Messages{
		"slice": []string{"min-5"},
	})

	// 普通元素指定为 asSlice
	v = New(ContinueAtError, message.NewPrinter(language.SimplifiedChinese))
	Min(5).Message("min-5").AsSlice().valid(v, "slice", "123456")
	a.Equal(v.Messages(), Messages{
		"slice": []string{"min-5"},
	})

	// ContinueAtError
	v = New(ContinueAtError, message.NewPrinter(language.SimplifiedChinese))
	Min(5).Message("min-5").AsSlice().valid(v, "slice", []int{1, 2, 6})
	a.Equal(v.Messages(), Messages{
		"slice[0]": []string{"min-5"},
		"slice[1]": []string{"min-5"},
	})

	// ExitAtError
	v = New(ExitAtError, message.NewPrinter(language.SimplifiedChinese))
	Min(5).Message("min-5").AsSlice().valid(v, "slice", []int{1, 2, 6})
	a.Equal(v.Messages(), Messages{
		"slice[0]": []string{"min-5"},
	})
}
