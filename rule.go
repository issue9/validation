// SPDX-License-Identifier: MIT

package validation

import (
	"reflect"
	"strconv"

	"github.com/issue9/localeutil"
	"golang.org/x/text/message"
)

// Rule 验证规则
type Rule struct {
	validator Validator
	asSlice   bool
	message   localeutil.LocaleStringer
}

// IfExpr 根据 if 条件选择不同的验证规则
type IfExpr struct {
	expr      bool
	ifRules   []*Rule
	elseRules []*Rule
}

// NewRule 返回 Rule 实例
func NewRule(validator Validator, key message.Reference, v ...any) *Rule {
	return &Rule{
		validator: validator,
		message:   localeutil.Phrase(key, v...),
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

func (rule *Rule) valid(v *Validation, name string, val any) bool {
	if !rule.asSlice {
		if !rule.validator.IsValid(val) {
			v.messages.Add(name, rule.message)
			return false
		}
		return true
	}

	// 以下 rule.asSlice == true

	rv := reflect.ValueOf(val)
	if kind := rv.Kind(); kind != reflect.Array && kind != reflect.Slice {
		ok := rule.validator.IsValid(val)
		if !ok {
			v.messages.Add(name, rule.message)
		}
		return ok
	}

	var sliceHasError bool
	for i := 0; i < rv.Len(); i++ {
		if !rule.validator.IsValid(rv.Index(i).Interface()) {
			v.messages.Add(name+"["+strconv.Itoa(i)+"]", rule.message)
			sliceHasError = true

			if v.errHandling != ContinueAtError {
				return false
			}
		}
	}
	return !sliceHasError
}

// If 返回 IfExpr 表达式
func If(expr bool, rule ...*Rule) *IfExpr {
	return &IfExpr{
		expr:    expr,
		ifRules: rule,
	}
}

// Else 指定条件不成言的验证规则
//
// 调用多次，则以最后一次指定为准，如果最后一次为空，则取消 Else 的内容。
func (expr *IfExpr) Else(rule ...*Rule) *IfExpr {
	expr.elseRules = rule
	return expr
}

// Rules 返回当前表达式最后使用的验证规则
func (expr *IfExpr) Rules() []*Rule {
	if expr.expr {
		return expr.ifRules
	}
	return expr.elseRules
}
