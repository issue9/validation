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

type ErrorHandling int8

// Validation 验证器
type Validation struct {
	errHandling ErrorHandling
	messages    Messages
	p           *message.Printer
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
// name 表示当前字段的名称，当验证出错时，以此值作为名称返回给用户；
// rules 表示验证的规则，按顺序依次验证。
func (v *Validation) NewField(val any, name string, rules ...*Rule) *Validation {
	if !v.messages.Empty() && v.errHandling == ExitAtError {
		return v
	}

	for _, rule := range rules {
		if !rule.valid(v, name, val) && v.errHandling != ContinueAtError {
			return v
		}
	}
	return v
}

// Messages 返回验证结果
func (v *Validation) Messages() Messages { return v.messages }
