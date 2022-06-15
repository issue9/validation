// SPDX-License-Identifier: MIT

// Package validator 提供各类验证器
package validator

type (
	// Validator 用于验证指定数据的合法性
	Validator interface {
		// IsValid 验证 v 是否符合当前的规则
		IsValid(v any) bool
	}

	// ValidateFunc 用于验证指定数据的合法性
	ValidateFunc func(any) bool
)

// IsValid 将当前函数作为 Validator 使用
func (f ValidateFunc) IsValid(v any) bool { return f(v) }

func And(v ...Validator) Validator {
	return ValidateFunc(func(a any) bool {
		for _, validator := range v {
			if !validator.IsValid(a) {
				return false
			}
		}
		return true
	})
}

func Or(v ...Validator) Validator {
	return ValidateFunc(func(a any) bool {
		for _, validator := range v {
			if validator.IsValid(a) {
				return true
			}
		}
		return false
	})
}
