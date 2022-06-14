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

type Validation struct {
	errHandling ErrorHandling
	messages    Messages
}

// New 返回 Validation 对象
//
// cap 表示初始的 Messages 容量大小；
func New(errHandling ErrorHandling, cap int) *Validation {
	return &Validation{
		errHandling: errHandling,
		messages:    make(Messages, cap),
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

// When 只有满足 cond 才执行 f 中的验证
//
// f 中的 v 即为当前对象；
func (v *Validation) When(cond bool, f func(v *Validation)) *Validation {
	if cond {
		f(v)
	}
	return v
}

// Messages 返回验证结果
func (v *Validation) Messages() Messages { return v.messages }

// LocaleMessages 返回本地化的验证结果
func (v *Validation) LocaleMessages(p *message.Printer) LocaleMessages { return Locale(v.messages, p) }
