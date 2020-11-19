/*
@Time : 2020/9/21 20:17
@Author : ZhouHui2
@File : 707
@Software: GoLand
*/
package linkedList

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList { // 头节点是伪的
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	p := this.Next
	i := 0
	for p.Next != nil {
		if i == index {
			return p.Val
		} else {
			p = p.Next
			i++
		}
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	head := this.Next
	p := &MyLinkedList{
		Val:  val,
		Next: head,
	}
	head = p
	this.Next = head
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	head := this.Next
	p := head
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &MyLinkedList{
		Val:  val,
		Next: nil,
	}
	this.Next = head
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
	} else {
		p := this.Next
		i := 1
		for p != nil {
			if i == index { //找到了位置,插入一个节点
				t1 := p.Next
				p = &MyLinkedList{
					Val:  val,
					Next: t1,
				}
				break
			} else if p.Next != nil {
				p = p.Next
				i++
			}
		}

	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this = this.Next
	}
	p := this
	i := 0
	for p.Next != nil {
		if i == index { //找到了位置
			p.Next = p.Next.Next
		} else {
			p = p.Next
			i++
		}
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
