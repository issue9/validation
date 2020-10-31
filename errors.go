// SPDX-License-Identifier: MIT

package validation

// Errors 表示一组错误信息的集合
//
// 键名查询参数名称，键值则为在解析和验证过种中返回的错误信息。
type Errors map[string][]string

// Add 为查询参数 key 添加一条新的错误信息
func (err Errors) Add(key string, val ...string) {
	if len(val) == 0 {
		panic("参数 val 必须指定")
	}

	err[key] = append(err[key], val...)
}

// Set 将查询参数 key 的错误信息改为 val
func (err Errors) Set(key string, val ...string) {
	if len(val) == 0 {
		panic("参数 val 必须指定")
	}

	err[key] = val
}
