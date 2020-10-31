// SPDX-License-Identifier: MIT

package validation

// Ruler 验证规则需要实现的接口
type Ruler interface {
	// 验证 v 是否符合当前的规则
	//
	// 如果不符合则返回具体的说明，否则返回空值。
	Validate(v interface{}) string
}

// RuleFunc 验证函数的签名
type RuleFunc func(v interface{}) string

// Validate 实现 Ruler
func (f RuleFunc) Validate(v interface{}) string {
	return f(v)
}
