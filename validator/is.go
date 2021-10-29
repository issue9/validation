// SPDX-License-Identifier: MIT

package validator

import (
	"github.com/issue9/validation"
	"github.com/issue9/validation/is"
)

// is 包中的简单封闭
var (
	GB32100  = validation.ValidateFunc(is.GB32100)
	GB11643  = validation.ValidateFunc(is.GB11643)
	HexColor = validation.ValidateFunc(is.HexColor)
	BankCard = validation.ValidateFunc(is.BankCard)
	ISBN     = validation.ValidateFunc(is.ISBN)
	URL      = validation.ValidateFunc(is.URL)
	IP       = validation.ValidateFunc(is.IP)
	IP4      = validation.ValidateFunc(is.IP4)
	IP6      = validation.ValidateFunc(is.IP6)
	Email    = validation.ValidateFunc(is.Email)

	CNPhone  = validation.ValidateFunc(is.CNPhone)
	CNMobile = validation.ValidateFunc(is.CNMobile)
	CNTel    = validation.ValidateFunc(is.CNTel)
)
