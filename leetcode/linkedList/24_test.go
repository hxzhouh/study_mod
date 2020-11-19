/*
@Time : 2020/9/15 11:15
@Author : ZhouHui2
@File : 24_test.go
@Software: GoLand
*/
package linkedList

import (
	"fmt"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	r := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	r = swapPairs(r)

	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
