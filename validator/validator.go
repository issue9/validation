// SPDX-License-Identifier: MIT

// Package validator 提供各类验证器
package validator

import (
	"regexp"

	"github.com/issue9/validation"
	"github.com/issue9/validation/is"
)

// Match 定义正则匹配的验证规则
func Match(exp *regexp.Regexp) validation.ValidateFunc {
	return validation.ValidateFunc(func(v interface{}) bool {
		return is.Match(exp, v)
	})
}

// Required 判断值是否必须为非空的规则
//
// skipNil 表示当前值为指针时，如果指向 nil，是否跳过非空检测规则。
// 如果 skipNil 为 false，则 nil 被当作空值处理。
//
// 具体判断规则可参考 github.com/issue9/is.Empty
func Required(skipNil bool) validation.ValidateFunc {
	return validation.ValidateFunc(func(v interface{}) bool {
		if skipNil && v == nil {
			return true
		}
		return !is.Empty(v, false)
	})
}
