package linkedList

import (
	"testing"
)

func Test_reversePrint(t *testing.T) {
	r := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	reversePrint(r)
}
