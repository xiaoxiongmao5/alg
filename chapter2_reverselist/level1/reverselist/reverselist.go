package reverselist

import (
	slink "xj/chapter2_reverselist/level1/single_linklist"
	// "fmt"
)



// 使用辅助头结点反转
func ReverselistByHead[T any](listNode *slink.LinkNode[T]) *slink.LinkNode[T]{
	if listNode == nil || listNode.Next == nil {
		return listNode
	}
	newHead := &slink.LinkNode[T]{}
	// 第一个有值结点
	tmp := listNode.Next
	for {
		next := tmp.Next
		// 下面两步将tmp结点挂在到新链表的首位
		tmp.Next = newHead.Next
		newHead.Next = tmp
		if next == nil {
			break
		}
		tmp = next
	}
	return newHead
}

// 自身反转
func ReverselistBySelf[T any](listNode *slink.LinkNode[T]) *slink.LinkNode[T] {
	if listNode == nil || listNode.Next == nil {
		return listNode
	}
	var newList *slink.LinkNode[T]
	tmp := listNode
	for {
		next := tmp.Next
		tmp.Next = newList
		newList = tmp
		if next == nil {
			break
		}
		tmp = next
	}
	return newList
}

// 使用递归反转 1->2->3
func ReverseListByRecursion[T any](listNode *slink.LinkNode[T]) *slink.LinkNode[T]{
	if listNode == nil || listNode.Next == nil {
		return listNode
	}
	tmp := listNode
	// newList将是该链表的尾节点
	newList := ReverseListByRecursion(tmp.Next)
	// 将tmp和它的下一个交换（eg: 2->3 ==> 3->2）
	tmp.Next.Next = tmp
	tmp.Next = nil
	return newList
}