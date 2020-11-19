package linkedList

import (
	"fmt"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
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
	r = removeNthFromEnd(r, 2)

	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
