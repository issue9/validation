// SPDX-License-Identifier: MIT

package validator

import (
	"reflect"

	"github.com/issue9/sliceutil"
	"golang.org/x/text/message"

	"github.com/issue9/validation"
)

// InValidator 判断一个值是否包含在指定元素中的验证器
type InValidator struct {
	not      bool
	elements []interface{}
}

// In 声明枚举类型的验证规则
//
// 要求验证的值必须包含在 element 元素中，如果不存在，则返回 msg 的内容。
func In(element ...interface{}) validation.Validator {
	return newInValidator(false, element...)
}

// NotIn 声明不在枚举中的验证规则
func NotIn(element ...interface{}) validation.Validator {
	return newInValidator(true, element...)
}

func newInValidator(not bool, element ...interface{}) *InValidator {
	return &InValidator{
		not:      not,
		elements: element,
	}
}

// Message 关联错误信息并返回 Message 实例
func (in *InValidator) Message(key message.Reference, v ...interface{}) *validation.Rule {
	return validation.NewRule(in, key, v...)
}

// IsValid 实现 validation.Validator
func (in *InValidator) IsValid(v interface{}) bool {
	isIn := sliceutil.Count(in.elements, func(i int) bool {
		elem := in.elements[i]
		elemType := reflect.TypeOf(elem)

		rv := reflect.ValueOf(v)
		if rv.Type().ConvertibleTo(elemType) && rv.Convert(elemType).Interface() == elem {
			return true
		}
		return reflect.DeepEqual(v, elem)
	}) > 0

	return (!in.not && isIn) || (in.not && !isIn)
}
