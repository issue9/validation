// SPDX-License-Identifier: MIT

package validation

import (
	"github.com/issue9/localeutil"
	"golang.org/x/text/message"
)

// LocaleMessages 表示一组错误信息的集合
//
// 键名查询参数名称，键值则为在解析和验证过种中返回的错误信息。
type LocaleMessages = MessagesOf[string]

// Messages 一组未本地化的消息集合
//
// 键名查询参数名称，键值则为在解析和验证过种中返回的错误信息。
type Messages = MessagesOf[localeutil.LocaleStringer]

type MessagesOf[T any] map[string][]T

// Add 为查询参数 key 添加一条新的错误信息
func (msg MessagesOf[T]) Add(key string, val ...T) {
	if len(val) == 0 {
		panic("参数 val 必须指定")
	}
	msg[key] = append(msg[key], val...)
}

// Set 将查询参数 key 的错误信息改为 val
func (msg MessagesOf[T]) Set(key string, val ...T) {
	if len(val) == 0 {
		panic("参数 val 必须指定")
	}
	msg[key] = val
}

func (msg MessagesOf[T]) Empty() bool { return len(msg) == 0 }

// Merge 将另一个 Messages 内容合并到当前实例
func (msg MessagesOf[T]) Merge(m MessagesOf[T]) {
	for key, mm := range m {
		msg.Add(key, mm...)
	}
}

func Locale(msg Messages, p *message.Printer) LocaleMessages {
	lm := make(LocaleMessages, len(msg))
	for k, v := range msg {
		msgs := make([]string, 0, len(v))
		for _, ls := range v {
			msgs = append(msgs, ls.LocaleString(p))
		}
		lm[k] = msgs
	}
	return lm
}
