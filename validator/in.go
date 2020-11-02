// SPDX-License-Identifier: MIT

package validator

import (
	"reflect"

	"github.com/issue9/sliceutil"

	"github.com/issue9/validation"
)

// In 声明枚举类型的验证规则
//
// 要求验证的值必须包含在 element 元素中，如果不存在，则返回 msg 的内容。
func In(element ...interface{}) validation.ValidateFunc {
	return isIn(false, element...)
}

// NotIn 声明不在枚举中的验证规则
func NotIn(element ...interface{}) validation.ValidateFunc {
	return isIn(true, element...)
}

func isIn(not bool, element ...interface{}) validation.ValidateFunc {
	return validation.ValidateFunc(func(v interface{}) bool {
		in := sliceutil.Count(element, func(i int) bool {
			elem := element[i]
			elemType := reflect.TypeOf(elem)

			rv := reflect.ValueOf(v)
			if rv.Type().ConvertibleTo(elemType) && rv.Convert(elemType).Interface() == elem {
				return true
			}
			return reflect.DeepEqual(v, elem)
		}) > 0

		return (!not && in) || (not && !in)
	})
}
