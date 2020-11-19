package linkedList

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
func deleteDuplicates_82(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		for head != nil && head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		return deleteDuplicates_82(head.Next)
	} else {
		head.Next = deleteDuplicates_82(head.Next)
		return head
	}
}
