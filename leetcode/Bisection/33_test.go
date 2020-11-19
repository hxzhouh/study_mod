package Bisection

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	//a := []int{4, 5, 6, 7, 0, 1, 2}
	a := []int{5, 1, 3}
	fmt.Println(search(a, 3))
}
