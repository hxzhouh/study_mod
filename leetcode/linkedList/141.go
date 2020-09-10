/*
@Time : 2020/9/10 11:46
@Author : ZhouHui2
@File : 141.go
@Software: GoLand
*/
package linkedList

func hasCycle(head *ListNode) bool {
	if head ==nil {
		return false
	}
	fast := head.Next
	for fast != nil && head != nil && fast.Next != nil {
		if fast == head {
			return true
		}
		fast= fast.Next.Next
		head=head.Next
	}
	return false
}
