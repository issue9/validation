// SPDX-License-Identifier: MIT

package validation_test

import (
	"testing"

	"github.com/issue9/assert/v2"

	"github.com/issue9/validation"
	"github.com/issue9/validation/validator"
)

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
