package test

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	a := []int{6, 3, 4, 5, 2, 1}
	//a := []int{1}
	fmt.Println("source:", a)
	fmt.Println("-------------------")
	selectionSort(a)
	fmt.Println("after:", a)
}

func BenchmarkBubbleSort(b *testing.B) {
	a := []int{6, 3, 4, 5, 2, 1}
	for i := 0; i < b.N; i++ {
		selectionSort(a)
	}
}
func BenchmarkMergeSort(b *testing.B) {
	a := []int{6, 3, 4, 5, 2, 1}
	for i := 0; i < b.N; i++ {
		MergeSort(a)
	}
}
