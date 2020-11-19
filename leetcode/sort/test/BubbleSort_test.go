/*
@Time : 2020/9/11 9:58
@Author : ZhouHui2
@File : BubbleSort_test.go
@Software: GoLand
*/
package test

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{4, 5, 6, 3, 2, 1}
	fmt.Println("source:", a)
	fmt.Println("-------------------")
	BubbleSort(a)
	fmt.Println("after:", a)
}
