package test

import (
	"fmt"
	"testing"
)

func TestQuicksort(t *testing.T) {
	a := []int{3, 2, 4, 6, 5, 1}
	fmt.Println("source:", a)
	fmt.Println("-------------------")
	a = Quicksort(a)
	fmt.Println("after:", a)
}
