/*
@Time : 2020/9/11 17:26
@Author : ZhouHui2
@File : InsertionSort_test.go
@Software: GoLand
*/
package test

import (
	"fmt"
	"testing"
)

func Test_insertionSort(t *testing.T) {
	a := []int{6, 5, 4, 3, 2, 1}
	fmt.Println("source:", a)
	fmt.Println("-------------------")
	insertionSort(a)
	fmt.Println("after:", a)
}
