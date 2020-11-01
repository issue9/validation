// SPDX-License-Identifier: MIT

package validation

import "golang.org/x/text/message"

// Validator 用于验证指定数据的合法性
type Validator interface {
	// 验证 v 是否符合当前的规则
	IsValid(v interface{}) bool
}

// ValidateFunc 用于验证指定数据的合法性
type ValidateFunc func(interface{}) bool

// Rule 验证规则需要实现的接口
type Rule struct {
	validator Validator

	key    message.Reference
	values []interface{}
}

// IsValid 将当前函数作为 Validator 使用
func (f ValidateFunc) IsValid(v interface{}) bool {
	return f(v)
}

// Message 将当前函数转换成 Message 实例
func (f ValidateFunc) Message(key message.Reference, v ...interface{}) *Rule {
	return NewRule(f, key, v...)
}

// NewRule 返回 Rule 实例
func NewRule(validator Validator, key message.Reference, v ...interface{}) *Rule {
	return &Rule{
		validator: validator,
		key:       key,
		values:    v,
	}
}

func (rule *Rule) message(p *message.Printer) string {
	return p.Sprintf(rule.key, rule.values...)
}

func (rule *Rule) isValid(v interface{}) bool {
	// TODO
	return rule.validator.IsValid(v)
}
