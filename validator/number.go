// SPDX-License-Identifier: MIT

package validator

import (
	"math"
	"reflect"

	"github.com/issue9/validation"
)

// Range 声明判断数值大小的验证规则
//
// 只能验证类型为 int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、float32 和 float64 类型的值。
//
// min 和 max 可以分别采用 math.Inf(-1) 和 math.Inf(1) 表示其最大的值范围。
func Range(min, max float64) validation.ValidateFunc {
	if max < min {
		panic("max 必须大于等于 min")
	}

	return validation.ValidateFunc(func(v interface{}) bool {
		var val float64
		switch v.(type) {
		case int, int8, int16, int32, int64:
			val = float64(reflect.ValueOf(v).Int())
		case uint, uint8, uint16, uint32, uint64:
			val = float64(reflect.ValueOf(v).Uint())
		case float32, float64:
			val = reflect.ValueOf(v).Float()
		default:
			return false
		}

		return val >= min && val <= max
	})
}

// Min 声明判断数值不小于 min 的验证规则
func Min(min float64) validation.ValidateFunc {
	return Range(min, math.Inf(1))
}

// Max 声明判断数值不大于 max 的验证规则
func Max(max float64) validation.ValidateFunc {
	return Range(math.Inf(-1), max)
}
