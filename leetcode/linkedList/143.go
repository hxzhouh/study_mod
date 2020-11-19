package linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {

	fast, slow := head, head

	for fast != nil {
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 翻转 slow
	slow = reverseList_26(slow)
	result := head

	for slow != nil && result != nil && slow != result {
		next := slow.Next
		if slow == result.Next { //边界
			slow.Next = nil
		} else {
			slow.Next = result.Next
		}
		result.Next = slow
		result = result.Next.Next
		slow = next
	}
}
