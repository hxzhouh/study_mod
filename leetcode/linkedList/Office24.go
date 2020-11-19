package linkedList

//定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
func reverseList(head *ListNode) *ListNode {
	var result *ListNode

	for head != nil {
		temp := head
		head = head.Next
		temp.Next = result
		result = temp

	}
	return result
}
