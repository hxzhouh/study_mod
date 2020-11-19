/*
@Time : 2020/9/10 15:27
@Author : ZhouHui2
@File : Stack
@Software: GoLand
*/
package LinkedListStack

import "github.com/hxzhouh/study_mod/leetcode/linkedList"

type Stack struct {
	count int
	Head  *linkedList.ListNode // 栈顶元素
	Tail  *linkedList.ListNode //栈底元素
}

func InitNewStack() *Stack {
	return &Stack{}
}
func (t *Stack) Pop() (int, error) {

	return 0, nil
}

func (t *Stack) Push(value int) error {
	node := &linkedList.ListNode{
		Val:  value,
		Next: nil,
	}
	if t.Tail == nil { //栈是空的
		t.count = 1
		t.Head = node
		t.Tail = node
	} else {
		t.count += 1
		t.Head.Next = node
		t.Head = t.Head.Next
	}

	return nil
}
func (t *Stack) IsExist(value int) (bool, error) {
	temp := t.Tail
	for temp.Next != nil {
		if temp.Val == value {
			return true, nil
		}
		temp = temp.Next
	}
	return false, nil
}
