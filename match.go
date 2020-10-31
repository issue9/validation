// SPDX-License-Identifier: MIT

package validation

import "regexp"

// Match 定义正则匹配的验证规则
func Match(exp *regexp.Regexp) ValidateFunc {
	return ValidateFunc(func(v interface{}) bool {
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
