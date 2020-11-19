package test

import (
	"fmt"
	"testing"
)

func TestBisection(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(Bisection(a, 8))
}
