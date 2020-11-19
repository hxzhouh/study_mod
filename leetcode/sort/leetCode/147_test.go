package leetCode

import (
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	r := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	insertionSortList(r)
}
