package linkedList

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
func deleteDuplicates_82(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	q := head
	var v int
	for p.Next != nil {
		if p.Val == p.Next.Val {
			v = p.Val
			p.Next = p.Next.Next
		} else if p.Val == v {
			if p.Next != nil {
				if p.Next.Next != nil {
					p = p.Next.Next
				}
			}
		} else {
			if p.Next != nil {
				p = p.Next
				q.Next = p
				q = q.Next
			}
		}
	}
	return head
}
