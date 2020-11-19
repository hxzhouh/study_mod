package linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	temp := result
	for {
		if l1 == nil || l2 == nil {
			break
		}
		if l1.Val > l2.Val {
			temp.Next = l2
			l2 = l2.Next
			temp = temp.Next
		} else {
			temp.Next = l1
			l1 = l1.Next
			temp = temp.Next
		}
	}
	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}
	return result.Next
}
