// SPDX-License-Identifier: MIT

package validator

import (
	"reflect"

	"github.com/issue9/validation"
)

// MinLength 声明判断内容长度不小于 min 的验证规则
func MinLength(min int64) validation.ValidateFunc {
	return Length(min, -1)
}

// MaxLength 声明判断内容长度不大于 max 的验证规则
func MaxLength(max int64) validation.ValidateFunc {
	return Length(-1, max)
}

// Length 声明判断内容长度的验证规则
//
// 如果 min 和 max 有值为 -1，表示忽略该值的比较，都为 -1 表示不限制长度。
//
// 只能验证类型为 string、Map、Slice 和 Array 的数据。
func Length(min, max int64) validation.ValidateFunc {
	if min > 0 && max > 0 && min > max {
		panic("max 必须大于 min")
	}

	return validation.ValidateFunc(func(v interface{}) bool {
		if min < 0 && max < 0 {
			return true
		}

		var l int64
		switch vv := v.(type) {
		case string:
			l = int64(len(vv))
		default:
			rv := reflect.ValueOf(v)
			switch rv.Kind() {
			case reflect.Array, reflect.Map, reflect.Slice:
				l = int64(rv.Len())
			default:
				return false
			}
		}

		if min < 0 {
			return l <= max // min 已经 i<=0，那么 max 必定 >=0
		}
		if max < 0 {
			return l >= min
		}
		return l >= min && l <= max
	})
}
