package ring_singlelink

import (
	"fmt"
)

type Cat struct {
	Id int
	Next *Cat
}

// 找到position的前一个节点返回
func FindLinkNode(headNode *Cat, position int) (bool, *Cat){
	preNode := headNode
	lastNode := headNode
	for {
		lastNode = lastNode.Next
		if lastNode.Next == headNode {
			break
		}
	}
	len := 1
	flag := false
	for {
		if len == position {
			flag = true
			break
		}
		len++
		preNode = lastNode
		lastNode = lastNode.Next
		if lastNode.Next == headNode {
			break
		}
	}
	if !flag && position==0 {
		return true, preNode	
	}
	return flag, lastNode
}

func AddAfter(node *Cat, newNode *Cat) {
	newNode.Next = node.Next
	node.Next = newNode
}

func DelNode(node *Cat) {
	if node.Next == node {
		node.Next = nil
		return
	}
	if node.Next != nil {
		node.Next = node.Next.Next
	}
}

// position: 0 最后一位；1 第一位（是新的headNode）；没找到位置时添加到最后一位；
func AddLinkNode(headNode *Cat, newNode *Cat, position int) *Cat{
	// 若该链表是否为空,自己指向自己
	if headNode.Next == nil {
		headNode.Id = newNode.Id
		headNode.Next = headNode
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
	if headNode.Next == nil {
		return headNode
	}
	find, tmpNode := FindLinkNode(headNode, position)
	if find {
		DelNode(tmpNode)
	}
	if !find && position == 0 {
		DelNode(tmpNode)
	}
	if position == 1 {
		return headNode.Next
	}
	return headNode
}

func ShowLink(headNode *Cat) {
	if headNode.Next == nil {
		fmt.Println("link empty")
		fmt.Println()
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