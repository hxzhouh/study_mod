package linkedList

import (
	"fmt"
	"testing"
)

func Test_reverseList(t *testing.T) {
	array := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	r := reverseList(array)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
