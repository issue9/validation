// SPDX-License-Identifier: MIT

// Package validation 数据验证相关功能
package validation

// 当验证出错时的几种可用处理方式
const (
	ContinueAtError  ErrorHandling = iota // 碰到错误不中断验证
	ExitAtError                           // 碰到错误中断验证
	ExitFieldAtError                      // 碰到错误中断当前字段的验证
)

// ErrorHandling 当验证出错时的处理方式
type ErrorHandling int8

// Validation 验证器
type Validation struct {
	ErrorHandling ErrorHandling
	errors        Errors
}

// NewField 验证新的字段
func (v *Validation) NewField(val interface{}, name string, rules ...Ruler) *Validation {
	if len(v.errors) > 0 && v.ErrorHandling == ExitAtError {
		return v
	}

	for _, rule := range rules {
		if msg := rule.Validate(val); msg != "" {
			v.errors.Add(name, msg)

			if v.ErrorHandling != ContinueAtError {
				return v
			}
		}
	}

	if len(v.errors[name]) > 0 { // 当前验证规则有错，则不验证子元素。
		return v
	}

	if vv, ok := val.(Validator); ok {
		if errors := vv.Validate(); len(errors) > 0 {
			for key, vals := range errors {
				v.errors.Add(name+key, vals...)
			}
		}
	}

	return v
}

// Result 返回验证结果
func (v *Validation) Result() Errors {
	return v.errors
}
