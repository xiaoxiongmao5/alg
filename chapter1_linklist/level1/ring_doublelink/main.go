package ring_doublelink

import (
	"fmt"
)

type Cat struct {
	Id int
	Pre *Cat
	Next *Cat
}

// 返回当前
func FindLinkNode(headNode *Cat, position int) (bool, *Cat) {
	tmpNode := headNode
	len := 1
	find := false
	for {
		// 找到添加位置了
		if len == position {
			find = true
			break
		}
		if tmpNode.Next == headNode {
			break
		}
		len++
		tmpNode = tmpNode.Next
	}
	if position == 0 && headNode.Next != nil {
		find = true
	}
	return find, tmpNode
}

func AddBefore(node *Cat, newNode *Cat){
	node.Pre.Next = newNode
	newNode.Pre = node.Pre
	newNode.Next = node
	node.Pre = newNode
}

func AddAfter(node *Cat, newNode *Cat){
	newNode.Pre = node
	newNode.Next = node.Next
	node.Next.Pre = newNode
	node.Next = newNode
}

func RemoveCur(node *Cat) {
	if node == node.Next {
		node.Next = nil
		node.Pre = nil
	}else{
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
	}
}

// position: 0 最后一位；1 第一位（是新的headNode）；没找到位置时添加到最后一位；
func AddLinkNode(headNode *Cat, newNode *Cat, position int) *Cat{

	// 若该链表是否为空,自己指向自己
	if headNode.Next == nil {
		headNode.Id = newNode.Id
		headNode.Next = headNode
		headNode.Pre = headNode
		return headNode
	}
	_, node := FindLinkNode(headNode, position)
	if position == 0 {
		AddAfter(node, newNode)
	} else {
		AddBefore(node, newNode)
	}
	var ret = headNode
	if position == 1 {
		ret = headNode.Pre
	}
	return ret
}

// position: 0 最后一位； 1第一位（headNode向下移动一位）； 没找到位置时不删除
func DelLinkNode(headNode *Cat, position int) *Cat{
	if headNode.Next == nil {
		return headNode
	}
	find, node := FindLinkNode(headNode, position)
	if !find {
		return headNode
	}
	RemoveCur(node)
	var ret = headNode
	if position == 1 && headNode.Next != nil{
		ret = headNode.Next
	}
	return ret
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

func ShowLink2(headNode *Cat) {
	if headNode.Next == nil {
		fmt.Println("link empty")
		fmt.Println()
		return
	}
	tmpNode := headNode
	for {
		fmt.Printf("Id=%d ==>\n", tmpNode.Pre.Id)
		tmpNode = tmpNode.Pre
		if tmpNode == headNode {
			break
		}
	}
	fmt.Println()
}