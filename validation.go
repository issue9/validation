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

	parentName string // 用于保存进入子元素的 ValidateFields 方法时，保存上一级的名称。
	separator  string // 用于指定上下级元素名称之间的分隔符。
}

// FieldsValidator 验证子项接口
//
// 一般用在自定义类型上，用于验证自身的子项数据。
//
// 凡实现此接口的对象，在 NewField 中会自动调用此接口的方法进行额外验证。
type FieldsValidator interface {
	ValidateFields(*Validation)
}

// New 返回 Validation 对象
//
// separator 用于指定字段名称上下级元素名称之间的连接符。比如在返回 xml 元素时，
// 可能会采用 root/element 的格式表示上下级，此时 separator 应设置为 /。
// 而在 json 中，可能会被转换成 root.element 的格式。
func New(errHandling ErrorHandling, p *message.Printer, separator string) *Validation {
	return &Validation{
		errHandling: errHandling,
		messages:    Messages{},
		p:           p,
		separator:   separator,
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

	var hasError bool
	for _, rule := range rules {
		if !rule.valid(v, v.joinName(name), val) {
			if v.errHandling != ContinueAtError {
				return v
			}
			hasError = true
		}
	}

	if hasError { // 当前验证规则有错，则不验证子元素。
		return v
	}

	if vv, ok := val.(FieldsValidator); ok {
		v.parentName = v.joinName(name)
		vv.ValidateFields(v)

		if v.parentName == name {
			v.parentName = ""
		} else {
			v.parentName = v.parentName[:len(v.parentName)-len(name)-len(v.separator)]
		}
	}

	return v
}

func (v *Validation) joinName(name string) string {
	if v.parentName == "" {
		return name
	}
	return v.parentName + v.separator + name
}

// Messages 返回验证结果
func (v *Validation) Messages() Messages { return v.messages }
