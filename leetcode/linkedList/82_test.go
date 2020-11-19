package linkedList

import (
	"fmt"
	"testing"
)

func Test_deleteDuplicates_82(t *testing.T) {
	r := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
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
	r = deleteDuplicates_82(r)

	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
