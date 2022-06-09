// SPDX-License-Identifier: MIT

package validator

import (
	"github.com/issue9/sliceutil"

	"github.com/issue9/validation"
)

type inValidator[T comparable] struct {
	not      bool
	elements []T
}

// In 声明枚举类型的验证规则
//
// 要求验证的值必须包含在 element 元素中，如果不存在，则返回 msg 的内容。
func In[T comparable](element ...T) validation.Validator {
	return newInValidator(false, element...)
}

// NotIn 声明不在枚举中的验证规则
func NotIn[T comparable](element ...T) validation.Validator {
	return newInValidator(true, element...)
}

func newInValidator[T comparable](not bool, element ...T) *inValidator[T] {
	return &inValidator[T]{
		not:      not,
		elements: element,
	}
}

// IsValid 实现 validation.Validator
func (in *inValidator[T]) IsValid(v any) bool {
	isIn := sliceutil.Exists(in.elements, func(elem T) bool { return elem == v })
	return (!in.not && isIn) || (in.not && !isIn)
}
