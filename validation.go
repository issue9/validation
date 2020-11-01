// SPDX-License-Identifier: MIT

// Package validation 数据验证
package validation

import "golang.org/x/text/message"

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
	p           *message.Printer
}

// FieldsValidator 验证子项接口
//
// 一般用在自定义类型上，用于验证自身的子项数据。
//
// 凡实现此接口的对象，在 NewField 中会自动调用此接口的方法进行额外验证。
type FieldsValidator interface {
	ValidateFields(ErrorHandling, *message.Printer) Messages
}

// New 返回 Validation 对象
func New(errHandling ErrorHandling, p *message.Printer) *Validation {
	return &Validation{
		errHandling: errHandling,
		messages:    Messages{},
		p:           p,
	}
}

// NewField 验证新的字段
//
// val 表示需要被验证的值，如果是一个对象且需要验证子字段，那么让对象实现 FieldsValidator 接口，
// 则会自动调用该方法验证子项，将会将验证完的信息返回给当前的 Validation 实例；
// name 表示当前字段的名称；rules 表示验证的规则，按顺序依资验证。
func (v *Validation) NewField(val interface{}, name string, rules ...*Rule) *Validation {
	if len(v.messages) > 0 && v.errHandling == ExitAtError {
		return v
	}

	for _, rule := range rules {
		if !rule.valid(v, name, val) {
			if v.errHandling != ContinueAtError {
				return v
			}
		}
	}

	if len(v.messages[name]) > 0 { // 当前验证规则有错，则不验证子元素。
		return v
	}

	if vv, ok := val.(FieldsValidator); ok {
		if errors := vv.ValidateFields(v.errHandling, v.p); len(errors) > 0 {
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
