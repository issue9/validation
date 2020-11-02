// SPDX-License-Identifier: MIT

package validator

import (
	"math"
	"testing"

	"github.com/issue9/assert"
)

func TestRange(t *testing.T) {
	a := assert.New(t)

	a.Panic(func() {
		Range(100, 5)
	})

	r := Range(5, math.MaxInt16)
	a.True(r.IsValid(5))
	a.True(r.IsValid(5.1))
	a.True(r.IsValid(math.MaxInt8))
	a.False(r.IsValid(math.MaxInt32))
	a.False(r.IsValid(-1))
	a.False(r.IsValid(-1.1))
	a.False(r.IsValid("5"))

	r = Min(6)
	a.True(r.IsValid(6))
	a.True(r.IsValid(10))
	a.False(r.IsValid(5))

	r = Max(6)
	a.True(r.IsValid(6))
	a.False(r.IsValid(10))
	a.True(r.IsValid(5))
	a.True(r.IsValid(uint(5)))
}
