/*
@Time : 2020/9/15 11:02
@Author : ZhouHui2
@File : 24
@Software: GoLand
*/
package linkedList

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head

	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		n1 := temp.Next
		n2 := n1.Next
		next := n2.Next
		n2.Next = n1
		n1.Next = next
		temp.Next = n2
		temp = n1
	}
	return dummyHead.Next
}
