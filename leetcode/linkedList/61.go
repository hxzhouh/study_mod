/*
@Time : 2020/9/15 14:51
@Author : ZhouHui2
@File : 61
@Software: GoLand
*/
package linkedList

//61. 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if k == 0 {
		return head
	}
	//p := &ListNode{}
	p := head // 尾部节点
	length := 1
	for p.Next != nil {
		p = p.Next
		length++
	}
	end := k % length
	p.Next = head       // 成环
	tail := &ListNode{} // 新的尾部节点
	tail.Next = head
	//head = head.Next
	for i := 0; i < length-end; i++ {
		tail = tail.Next
		head = head.Next
	}
	tail.Next = nil
	return head
}
