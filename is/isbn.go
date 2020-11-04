// SPDX-License-Identifier: MIT

package is

import "bytes"

// ISBN 判断是否为合法的 ISBN 串号
//
// 可以同时判断 ISBN10 和 ISBN13
//
// 有关 ISBN 的算法及其它相关内容，可参照http://zh.wikipedia.org/wiki/%E5%9B%BD%E9%99%85%E6%A0%87%E5%87%86%E4%B9%A6%E5%8F%B7
func ISBN(val interface{}) bool {
	var result []byte

	switch v := val.(type) {
	case []byte:
		result = v
	case []rune:
		result = []byte(string(v))
	case string:
		result = []byte(v)
	default:
		return false
	}

	if bytes.IndexByte(result, '-') > -1 {
		result = eraseMinus(result)
	}

	switch len(result) {
	case 10:
		return isISBN10(result)
	case 13:
		return isISBN13(result)
	default:
		return false
	}
}

// ISBN10 判断是否为合法的 ISBN10
func ISBN10(val []byte) bool {
	if bytes.IndexByte(val, '-') == -1 {
		return isISBN10(val)
	}
	return isISBN10(eraseMinus(val))
}

// ISBN13 判断是否为合法的 ISBN13
func ISBN13(val []byte) bool {
	if bytes.IndexByte(val, '-') == -1 {
		return isISBN13(val)
	}
	return isISBN13(eraseMinus(val))
}

// isbn10 的校验位对应的值。
var isbn10Map = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'x', '0'}

func isISBN10(val []byte) bool {
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(val[i]-'0') * (10 - i)
	}

	if val[9] == 'X' {
		val[9] = 'x'
	}

	return isbn10Map[11-sum%11] == val[9]
}

func isISBN13(val []byte) bool {
	sum := 0
	for i := 0; i < 12; i += 2 {
		sum += int(val[i] - '0')
	}

	for i := 1; i < 12; i += 2 {
		sum += (int(val[i]-'0') * 3)
	}

	return (10 - sum%10) == int(val[12]-'0')
}

// 过滤减号(-)符号
func eraseMinus(val []byte) []byte {
	offset := 0
	for k, v := range val {
		if v == '-' {
			offset++
			continue
		}

		val[k-offset] = v
	}
	return val[:len(val)-offset]
}
