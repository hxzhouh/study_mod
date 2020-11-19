package linkedList

func reverseList_26(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
