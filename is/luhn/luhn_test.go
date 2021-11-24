// SPDX-License-Identifier: MIT

package luhn

import (
	"testing"

	"github.com/issue9/assert/v2"
)

func TestIsValid(t *testing.T) {
	a := assert.New(t, false)

	a.True(IsValid([]byte("6259650871772098")))
	a.True(IsValid([]byte("79927398713")))
	a.False(IsValid([]byte("79927398710")))
}

func TestGenerateWithPrefix(t *testing.T) {
	a := assert.New(t, false)

	a.Equal("6259650871772098", string(GenerateWithPrefix([]byte("625965087177209"))))
	a.True(IsValid(GenerateWithPrefix([]byte("625965087177209"))))
}
