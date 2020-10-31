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
	errHandling ErrorHandling
	messages    Messages
}

// Validator 验证对象接口
//
// 凡实现此接口的对象，在 NewField 中会自动调用此接口的方法进行额外验证。
type Validator interface {
	Validate(ErrorHandling) Messages
}

// New 返回 Validation 对象
func New(errHandling ErrorHandling) *Validation {
	return &Validation{
		errHandling: errHandling,
		messages:    Messages{},
	}
}

// NewField 验证新的字段
func (v *Validation) NewField(val interface{}, name string, rules ...Ruler) *Validation {
	if len(v.messages) > 0 && v.errHandling == ExitAtError {
		return v
	}

	for _, rule := range rules {
		if msg := rule.Validate(val); msg != "" {
			v.messages.Add(name, msg)

			if v.errHandling != ContinueAtError {
				return v
			}
		}
	}

	if len(v.messages[name]) > 0 { // 当前验证规则有错，则不验证子元素。
		return v
	}

	if vv, ok := val.(Validator); ok {
		if errors := vv.Validate(v.errHandling); len(errors) > 0 {
			for key, vals := range errors {
				v.messages.Add(name+key, vals...)
			}
		}
	}

	return v
}

// Messages 返回验证结果
func (v *Validation) Messages() Messages {
	return v.messages
}
