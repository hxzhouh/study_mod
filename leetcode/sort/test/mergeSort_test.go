package test

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	a := []int{6, 2, 4, 3, 5, 1}
	fmt.Println("source:", a)
	fmt.Println("-------------------")
	a = MergeSort(a)
	fmt.Println("after:", a)
}
