package leetCode

//https://leetcode-cn.com/problems/insertion-sort-list/
// 效率很低，不知道为啥
// 优化
/*
	// 优化部分： 如果比排序链表的最后一个大，直接插到末尾
	if end.Val < cur.Val {
		cur.Next = nil
		end.Next = cur
		end = cur
		cur = next
		continue
	}
*/
func insertionSortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	prev := head.Next // 从第二位开始插入
	for prev != nil { // 终止条件
		temp := head
		for temp != prev {
			if temp.Val > prev.Val { // 交换
				temp.Val += prev.Val
				prev.Val = temp.Val - prev.Val
				temp.Val = temp.Val - prev.Val
			}
			temp = temp.Next
		}
		prev = prev.Next
	}
	return head
}
