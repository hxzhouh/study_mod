/*
@Time : 2020/9/15 15:05
@Author : ZhouHui2
@File : 61_test.go
@Software: GoLand
*/
package linkedList

import (
	"fmt"
	"testing"
)

func Test_rotateRight(t *testing.T) {
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
	r1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	r = rotateRight(r1, 2)

	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
