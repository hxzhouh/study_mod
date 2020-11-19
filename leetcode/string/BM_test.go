package string

import (
	"fmt"
	"testing"
)

func Test_bm(t *testing.T) {
	a := []rune("abcdef")
	b := []rune("cde")

	fmt.Println(bm(a, b, len(a), len(b)))
}
