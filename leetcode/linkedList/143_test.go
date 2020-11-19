package linkedList

import (
	"fmt"
	"testing"
)

func Test_reorderList(t *testing.T) {
	test := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	reorderList(test)
	for test != nil {
		fmt.Println(test.Val)
		test = test.Next
	}
}
