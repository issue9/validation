// SPDX-License-Identifier: MIT

package validator

import (
	"math"

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

	return func(v any) bool {
		var val float64
		switch vv := v.(type) {
		case int:
			val = float64(vv)
		case int8:
			val = float64(vv)
		case int16:
			val = float64(vv)
		case int32:
			val = float64(vv)
		case int64:
			val = float64(vv)
		case uint:
			val = float64(vv)
		case uint8:
			val = float64(vv)
		case uint16:
			val = float64(vv)
		case uint32:
			val = float64(vv)
		case uint64:
			val = float64(vv)
		case float32:
			val = float64(vv)
		case float64:
			val = vv
		default:
			return false
		}

		return val >= min && val <= max
	}
}

// Min 声明判断数值不小于 min 的验证规则
func Min(min float64) validation.ValidateFunc { return Range(min, math.Inf(1)) }

// Max 声明判断数值不大于 max 的验证规则
func Max(max float64) validation.ValidateFunc { return Range(math.Inf(-1), max) }
