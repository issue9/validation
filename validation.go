// SPDX-License-Identifier: MIT

// Package validation 数据验证
package validation

import (
	"reflect"
	"strconv"

	"github.com/issue9/localeutil"
	"golang.org/x/text/message"

	"github.com/issue9/validation/validator"
)

// 当验证出错时的几种可用处理方式
const (
	ContinueAtError  ErrorHandling = iota // 碰到错误不中断验证
	ExitAtError                           // 碰到错误中断验证
	ExitFieldAtError                      // 碰到错误中断当前字段的其它规则验证
)

type (
	ErrorHandling int8

	Validation struct {
		errHandling ErrorHandling
		messages    Messages
	}

	Validator = validator.Validator

	ValidateFunc = validator.ValidateFunc

	// Rule 验证规则
	//
	// 这是对 Validator 的二次包装，保存着未本地化的错误信息，用以在验证失败之后返回给 Validation。
	Rule struct {
		validator Validator
		message   localeutil.LocaleStringer
	}
)

func NewRule(validator Validator, key message.Reference, v ...any) *Rule {
	return &Rule{
		validator: validator,
		message:   localeutil.Phrase(key, v...),
	}
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
		if rule.validator.IsValid(val) {
			continue
		}

		v.messages.Add(name, rule.message)
		if v.errHandling != ContinueAtError {
			break
		}
	}
	return v
}

// NewSliceField 验证数组字段
//
// 如果字段类型不是数组或是字符串，将直接返回错误。
func (v *Validation) NewSliceField(val any, name string, rules ...*Rule) *Validation {
	// TODO: 如果 go 支持泛型方法，那么可以将 val 固定在 []T

	rv := reflect.ValueOf(val)

	if kind := rv.Kind(); kind != reflect.Array && kind != reflect.Slice && kind != reflect.String {
		if v.errHandling != ContinueAtError {
			v.messages.Add(name, rules[0].message) // 非数组，取第一个规则的错误信息
			return v
		}
		for _, rule := range rules {
			v.messages.Add(name, rule.message)
		}
		return v
	}

	for i := 0; i < rv.Len(); i++ {
		for _, rule := range rules {
			if !rule.validator.IsValid(rv.Index(i).Interface()) {
				v.messages.Add(name+"["+strconv.Itoa(i)+"]", rule.message)
				if v.errHandling != ContinueAtError {
					return v
				}
			}
		}
	}

	return v
}

// NewMapField 验证 map 字段
//
// 如果字段类型不是 map，将直接返回错误。
func (v *Validation) NewMapField(val any, name string, rules ...*Rule) *Validation {
	// TODO: 如果 go 支持泛型方法，那么可以将 val 固定在 map[T]T

	rv := reflect.ValueOf(val)

	if kind := rv.Kind(); kind != reflect.Map {
		if v.errHandling != ContinueAtError {
			v.messages.Add(name, rules[0].message) // 非数组，取第一个规则的错误信息
			return v
		}
		for _, rule := range rules {
			v.messages.Add(name, rule.message)
		}
		return v
	}

	keys := rv.MapKeys()
	for i := 0; i < rv.Len(); i++ {
		key := keys[i]
		for _, rule := range rules {
			if !rule.validator.IsValid(rv.MapIndex(key).Interface()) {
				v.messages.Add(name+"["+key.String()+"]", rule.message)
				if v.errHandling != ContinueAtError {
					return v
				}
			}
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
