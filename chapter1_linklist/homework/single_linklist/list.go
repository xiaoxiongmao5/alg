package single_linklist

import (
	"fmt"
)

type LinkList[T any] struct {
	Head *LinkNode[T]
}

type LinkNode[T any] struct {
	Val T
	Next *LinkNode[T]
}

// 找到position前一个节点，没找到时返回最后一个节点
func (this *LinkList[T])FindLinkNode(position int) (bool, *LinkNode[T]){
	if this.Head.Next == nil {
		return false, this.Head
	}
	tmpNode := this.Head
	len := 0
	find := false
	for {
		if tmpNode.Next == nil {
			break
		}
		// 找到添加位置了
		if len + 1 == position {
			find = true
			break
		}
		len++
		tmpNode = tmpNode.Next
	}
	return find, tmpNode
}
// 添加元素到node之后
func (this *LinkList[T])AddAfter(node *LinkNode[T], newNode *LinkNode[T]) {
	newNode.Next = node.Next
	node.Next = newNode
}
// 删除node的下一个元素
func (this *LinkList[T])DelNodeAfter(node *LinkNode[T]) {
	if node.Next != nil {
		node.Next = node.Next.Next
	}
}
// 添加结点到最后一位
func (this *LinkList[T])AddLastNode(newNode *LinkNode[T]) {
	if this.Head == nil {
		this.Head.Next = newNode
		return
	}
	tmp := this.Head
	for {
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	tmp.Next = newNode
}
// 删除最后一个结点
func (this *LinkList[T])DelLastNode() {
	if this.Head.Next == nil {
		return
	}
	tmp, pre := this.Head, this.Head
	for {
		if tmp.Next == nil {
			break
		}
		pre = tmp
		tmp = tmp.Next
	}
	pre.Next = nil
}
// position: 0 最后一位；1 第一位；没找到位置时添加到最后一位；
func (this *LinkList[T])AddLinkNode(newNode *LinkNode[T], position int) {
	if position == 0 {
		this.AddLastNode(newNode)
		return
	}
	_, tmpNode := this.FindLinkNode(position)
	this.AddAfter(tmpNode, newNode)
}
// position: 0 最后一位； 1第一位； 没找到位置时不删除
func (this *LinkList[T])DelLinkNode(position int){
	if position == 0 {
		this.DelLastNode()
		return
	}
	find, tmpNode := this.FindLinkNode(position)
	if find {
		this.DelNodeAfter(tmpNode)
	}
}

func (this *LinkList[T])ShowLink() {
	if this.Head.Next == nil {
		fmt.Println("link empty")
		fmt.Println()
		return
	}
	tmpNode := this.Head
	for {
		fmt.Printf("Val=%v ==>\n", tmpNode.Next.Val)
		tmpNode = tmpNode.Next
		if tmpNode.Next == nil {
			break
		}
	}
	fmt.Println()
}

// 获得链表的长度
func (this *LinkList[T])GetListLen() int{
	if this.Head.Next == nil {
		return 0
	}
	len := 1
	tmp := this.Head.Next
	for {
		if tmp.Next == nil {
			break
		}
		len++
		tmp = tmp.Next
	}
	return len
}