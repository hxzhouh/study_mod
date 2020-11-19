package linkedList

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	//curr := head
	//length := 0
	//for curr != nil {
	//	curr = curr.Next
	//	length += 1
	//}
	//if length <= 1 {
	//	return nil
	//}
	//curr = dummy
	//length -= n
	//for length > 0 {
	//	length--
	//	curr = curr.Next
	//}
	//curr.Next = curr.Next.Next
	//return dummy.Next
	fast := dummy
	slow := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
