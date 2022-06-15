// SPDX-License-Identifier: MIT

package validator

import (
	"regexp"

	"github.com/issue9/validation/is"
)

// 对 is 包中的简单封装
var (
	GB32100  = ValidateFunc(is.GB32100)
	GB11643  = ValidateFunc(is.GB11643)
	HexColor = ValidateFunc(is.HexColor)
	BankCard = ValidateFunc(is.BankCard)
	ISBN     = ValidateFunc(is.ISBN)
	URL      = ValidateFunc(is.URL)
	IP       = ValidateFunc(is.IP)
	IP4      = ValidateFunc(is.IP4)
	IP6      = ValidateFunc(is.IP6)
	Email    = ValidateFunc(is.Email)

	CNPhone  = ValidateFunc(is.CNPhone)
	CNMobile = ValidateFunc(is.CNMobile)
	CNTel    = ValidateFunc(is.CNTel)
)

// Match 定义正则匹配的验证规则
func Match(exp *regexp.Regexp) ValidateFunc {
	return func(v any) bool {
		return is.Match(exp, v)
	}
}

// Required 判断值是否必须为非空的规则
//
// skipNil 表示当前值为指针时，如果指向 nil，是否跳过非空检测规则。
// 如果 skipNil 为 false，则 nil 被当作空值处理。
//
// 具体判断规则可参考 github.com/issue9/validation/is.Empty
func Required(skipNil bool) ValidateFunc {
	return func(v any) bool {
		if skipNil && v == nil {
			return true
		}
		return !is.Empty(v, false)
	}
}
