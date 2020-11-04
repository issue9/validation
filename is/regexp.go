// SPDX-License-Identifier: MIT

package is

import "regexp"

const (
	// 匹配大陆电话
	cnPhonePattern = `((\d{3,4})-?)?` + // 区号
		`\d{5,10}` + // 号码，95500等5位数的，7位，8位，以及400开头的10位数
		`(-\d{1,4})?` // 分机号，分机号的连接符号不能省略。

	// 匹配大陆手机号码
	cnMobilePattern = `(0|\+?86)?` + // 匹配 0,86,+86
		`(13[0-9]|` + // 130-139
		`14[4579]|` + // 144,145,147,149
		`15[0-9]|` + // 150-159
		`17[0-9]|` + // 170-179
		`18[0-9]|` + // 180-189
		`16[56]|` + // 165,1666
		`19[0126789])` + // 191,192,196,197,198,199
		`[0-9]{8}`

	// 匹配大陆手机号或是电话号码
	cnTelPattern = "(" + cnPhonePattern + ")|(" + cnMobilePattern + ")"

	// 匹配邮箱
	emailPattern = `[\w.-]+@[\w_-]+\w{1,}[\.\w-]+`

	// 匹配 IP4
	ip4Pattern = `((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)`

	// 匹配 IP6，参考以下网页内容：
	// http://blog.csdn.net/jiangfeng08/article/details/7642018
	ip6Pattern = `(([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|` +
		`(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|` +
		`(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|` +
		`(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
		`(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
		`(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
		`(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
		`(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))`

		// 同时匹配 IP4 和 IP6
	ipPattern = "(" + ip4Pattern + ")|(" + ip6Pattern + ")"

	// 匹配域名
	domainPattern = `[a-zA-Z0-9][a-zA-Z0-9_-]{0,62}(\.[a-zA-Z0-9][a-zA-Z0-9_-]{0,62})*(\.[a-zA-Z][a-zA-Z0-9]{0,10}){1}`

	// 匹配 URL
	urlPattern = `((https|http|ftp|rtsp|mms)?://)?` + // 协议
		`(([0-9a-zA-Z]+:)?[0-9a-zA-Z_-]+@)?` + // pwd:user@
		"(" + ipPattern + "|(" + domainPattern + "))" + // IP 或域名
		`(:\d{1,5})?` + // 端口
		`(/+[a-zA-Z0-9][a-zA-Z0-9_.-]*)*/*` + // path
		`(\?([a-zA-Z0-9_-]+(=.*&?)*)*)*` // query
)

var (
	email    = regexpCompile(emailPattern)
	ip4      = regexpCompile(ip4Pattern)
	ip6      = regexpCompile(ip6Pattern)
	ip       = regexpCompile(ipPattern)
	url      = regexpCompile(urlPattern)
	cnPhone  = regexpCompile(cnPhonePattern)
	cnMobile = regexpCompile(cnMobilePattern)
	cnTel    = regexpCompile(cnTelPattern)
)

func regexpCompile(str string) *regexp.Regexp {
	return regexp.MustCompile("^" + str + "$")
}

// Match 判断 val 是否能正确匹配 exp 中的正则表达式
//
// val 可以是[]byte, []rune, string类型。
func Match(exp *regexp.Regexp, val interface{}) bool {
	switch v := val.(type) {
	case []rune:
		return exp.MatchString(string(v))
	case []byte:
		return exp.Match(v)
	case string:
		return exp.MatchString(v)
	default:
		return false
	}
}

// CNPhone 验证中国大陆的电话号码
//
// 支持如下格式：
//  0578-12345678-1234
//  057812345678-1234
// 若存在分机号，则分机号的连接符不能省略。
func CNPhone(val interface{}) bool {
	return Match(cnPhone, val)
}

// CNMobile 验证中国大陆的手机号码
func CNMobile(val interface{}) bool {
	return Match(cnMobile, val)
}

// CNTel 验证手机和电话类型
func CNTel(val interface{}) bool {
	return Match(cnTel, val)
}

// URL 验证一个值是否标准的URL格式
//
// 支持 IP 和域名等格式
func URL(val interface{}) bool {
	return Match(url, val)
}

// IP 验证一个值是否为 IP
//
// 可验证 IP4 和 IP6
func IP(val interface{}) bool {
	return Match(ip, val)
}

// IP6 验证一个值是否为 IP6
func IP6(val interface{}) bool {
	return Match(ip6, val)
}

// IP4 验证一个值是滞为 IP4
func IP4(val interface{}) bool {
	return Match(ip4, val)
}

// Email 验证一个值是否匹配一个邮箱
func Email(val interface{}) bool {
	return Match(email, val)
}
