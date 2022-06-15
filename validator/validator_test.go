// SPDX-License-Identifier: MIT

package validator

import (
	"testing"

	"github.com/issue9/assert/v2"
)

func TestAnd(t *testing.T) {
	a := assert.New(t, false)

	v := And(In(1, 2, 3), NotIn(2, 3, 4))
	a.True(v.IsValid(1))
	a.False(v.IsValid(2))
	a.False(v.IsValid(-1))
	a.False(v.IsValid(100))
}

func TestOr(t *testing.T) {
	a := assert.New(t, false)

	v := Or(In(1, 2, 3), NotIn(2, 3, 4))
	a.True(v.IsValid(1))
	a.True(v.IsValid(2))
	a.False(v.IsValid(4))
	a.True(v.IsValid(-1))
	a.True(v.IsValid(100))
}
