/*
@Time : 2020/9/21 20:30
@Author : ZhouHui2
@File : 707_test.go
@Software: GoLand
*/
package linkedList

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	list := Constructor()
	list.AddAtHead(1)
	fmtMyLinkedList(list)
	list.AddAtTail(3)
	fmtMyLinkedList(list)
	list.AddAtIndex(1, 2)
	fmtMyLinkedList(list)
}
func fmtMyLinkedList(list MyLinkedList) {
	p := list.Next
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
	fmt.Println("====================================")
}
