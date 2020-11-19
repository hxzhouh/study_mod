package linkedList

func reversePrint(head *ListNode) []int {
	result := make([]int, 0)
	var prev, curr *ListNode
	curr = head //当前位置
	for curr != nil {
		next := curr.Next
		curr.Next = prev //翻转的关键
		prev = curr
		curr = next
	}
	for prev != nil {
		result = append(result, prev.Val)
		prev = prev.Next
	}
	return result
}
