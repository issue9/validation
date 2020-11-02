// SPDX-License-Identifier: MIT

package validator

import (
	"regexp"

	"github.com/issue9/validation"
)

// Match 定义正则匹配的验证规则
func Match(exp *regexp.Regexp) validation.ValidateFunc {
	return validation.ValidateFunc(func(v interface{}) bool {
		switch vv := v.(type) {
		case string:
			return exp.MatchString(vv)
		case []byte:
			return exp.Match(vv)
		case []rune:
			return exp.MatchString(string(vv))
		default:
			return false
		}
	})
}
