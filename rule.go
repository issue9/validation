// SPDX-License-Identifier: MIT

package validation

import (
	"reflect"
	"strconv"

	"golang.org/x/text/message"
)

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
	asSlice   bool

	// 错误提示信息
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

// AsSlice 以数组的形式验证数据
//
// 如果指定此属性，则所有 kind 值为 reflect.Array 和 reflect.Slice 的都将被当作数组处理，
// 包括 []byte 和 []rune 等。
// 其它类型继续以正常元素处理。
//
// 如果未指定此属性，则所有类型的元素都被当作一个值进行验证，即使是数组。
func (rule *Rule) AsSlice() *Rule {
	rule.asSlice = true
	return rule
}

func (rule *Rule) message(p *message.Printer) string {
	return p.Sprintf(rule.key, rule.values...)
}

func (rule *Rule) valid(v *Validation, name string, val interface{}) bool {
	if !rule.asSlice {
		if !rule.validator.IsValid(val) {
			v.messages.Add(name, rule.message(v.p))
			return false
		}
		return true
	}

	rv := reflect.ValueOf(val)
	if kind := rv.Kind(); kind != reflect.Array && kind != reflect.Slice {
		ok := rule.validator.IsValid(val)
		if !ok {
			v.messages.Add(name, rule.message(v.p))
		}
		return ok
	}

	var sliceHasError bool
	for i := 0; i < rv.Len(); i++ {
		if !rule.validator.IsValid(rv.Index(i).Interface()) {
			v.messages.Add(name+"["+strconv.Itoa(i)+"]", rule.message(v.p))
			sliceHasError = true

			if v.errHandling != ContinueAtError {
				return false
			}
		}
	}
	return !sliceHasError
}
