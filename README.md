# validation

[![Go](https://github.com/issue9/validation/workflows/Test/badge.svg)](https://github.com/issue9/validation/actions?query=workflow%3ATest)
[![Go Report Card](https://goreportcard.com/badge/github.com/issue9/validation)](https://goreportcard.com/report/github.com/issue9/validation)
![License](https://img.shields.io/github/license/issue9/validation)
[![Go version](https://img.shields.io/github/go-mod/go-version/issue9/validation)](https://golang.org)
[![codecov](https://codecov.io/gh/issue9/validation/branch/master/graph/badge.svg)](https://codecov.io/gh/issue9/validation)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/issue9/validation)](https://pkg.go.dev/github.com/issue9/validation)

数据验证

包含了以下类型的数据验证：

- GB11643 身份证验证和分析；
- GB32100 统一信用代码验证和分析；
- 银行卡卡号验证；
- ISBN 验证；
- IPv4/IPv6 验证；
- 中国区手机/电话号码验证；
- 域名验证；
- Email 验证；

```go
import (
    "github.com/issue9/validation"
    "github.com/issue9/validation/validator"
    "golang.org/x/text/language"
    "golang.org/x/text/message"
)

type Object {
    Age int
    Name string
}

o := &Object{}

p := message.NewPrinter(language.MustParse("cmn-Hans"))
v := validation.New(validation.ContinueAtError, p, ".")
messages := v.NewField(&o.Age, "age", validator.Min(18).Message("必须大于 18")).
    NewField(&o.Name, "name", validator.Required(false).Message("不能为空")).
    Messages()
```

## 本地化

本地化采用 golang.org/x/text 包

```go
import (
    "github.com/issue9/validation"
    "github.com/issue9/validation/validator"
    "golang.org/x/text/language"
    "golang.org/x/text/message"
    "golang.org/x/text/message/catalog"
)

type Object {
    Age int
    Name string
}

builder := catalog.NewBuilder()
builder.SetString(language.SimplifiedChinese, "lang", "chn")
builder.SetString(language.TraditionalChinese, "lang", "cht")

o := &Object{}

p := message.NewPrinter(language.SimplifiedChinese, message.Catalog(builder))
v := validation.New(validation.ContinueAtError, p, ".")
messages := v.NewField(&o.Age, "age", validator.Min(18).Message("lang")). // 根据 p 的不同，会输出不同内容
    NewField(&o.Name, "name", validator.Required(false).Message("不能为空")).
    Messages()
```

## 版权

本项目采用 [MIT](https://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。
