package ring_singlelink

import (
	"fmt"
)

type Cat struct {
	Id int
	Next *Cat
}

// 找到position前一个节点，没找到时返回最后一个节点
func FindLinkNode(headNode *Cat, position int) (bool, *Cat){
	if headNode.Next == nil {
		return false, headNode
	}
	tmp := headNode
	pre := headNode
	// 把pre置为链表最后一个
	for {
		pre = pre.Next
		if pre.Next == headNode {
			break
		}
	}
	len := 1
	find := false
	for {
		if len == position {
			find = true
			break
		}
		if tmp.Next == headNode {
			break
		}
		len++
		pre = pre.Next
		tmp = tmp.Next
	}
	if find {
		return find, pre
	}
	return find, tmp
}
// 在node后插入newNode
func AddAfter(node *Cat, newNode *Cat) {
	newNode.Next = node.Next
	node.Next = newNode
}
// 删除node的下一个元素
func DelNodeAfter(node *Cat) {
	if node.Next == node {
		node.Next = nil
		return
	}
	if node.Next != nil {
		node.Next = node.Next.Next
	}
}
// 添加结点到最后一位
func AddLastNode(headNode *Cat, newNode *Cat) {
	if headNode.Next == nil {
		headNode.Id = newNode.Id
		headNode.Next = headNode
		return
	}
	tmp := headNode
	for {
		if tmp.Next == headNode{
			break
		}
		tmp = tmp.Next
	}
	newNode.Next = tmp.Next
	tmp.Next = newNode
}
// 删除最后一个结点
func DelLastNode(headNode *Cat) {
	if headNode.Next == nil {
		return
	}
	// 只有一个节点自身循环
	if headNode.Next == headNode {
		headNode.Next = nil
		return
	}
	tmp := headNode
	pre := headNode
	for {
		if tmp.Next == headNode {
			break
		}
		pre = tmp
		tmp = tmp.Next
	}
	pre.Next = tmp.Next
}
// position: 0 最后一位；1 第一位（是新的headNode）；没找到位置时添加到最后一位；
func AddLinkNode(headNode *Cat, newNode *Cat, position int) *Cat{
	if position == 0 {
		AddLastNode(headNode, newNode)
		return headNode
	}
	_, tmpNode := FindLinkNode(headNode, position)
	AddAfter(tmpNode, newNode)
	if position == 1 {
		return newNode
	}
	return headNode
}
// position: 0 最后一位； 1第一位（headNode向下移动一位）； 没找到位置时不删除
func DelLinkNode(headNode *Cat, position int) *Cat{
	if position == 0 {
		DelLastNode(headNode)
		return headNode
	}
	find, tmpNode := FindLinkNode(headNode, position)
	if find {
		DelNodeAfter(tmpNode)
		if position == 1 {
			return headNode.Next
		}
	}
	return headNode
}

func ShowLink(headNode *Cat) {
	if headNode.Next == nil {
		fmt.Println("link empty")
		fmt.Println()
		return
	}
	if headNode.Next == headNode {
		fmt.Printf("Id=%d ==>\n", headNode.Id)
		return
	}
	tmpNode := headNode
	for {
		fmt.Printf("Id=%d ==>\n", tmpNode.Id)
		if tmpNode.Next == headNode {
			break
		}
		tmpNode = tmpNode.Next
	}
	fmt.Println()
}