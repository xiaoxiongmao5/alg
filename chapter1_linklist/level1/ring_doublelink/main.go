package ring_doublelink

import (
	"fmt"
)

type Cat struct {
	Id int
	Pre *Cat
	Next *Cat
}

// 找到position的当前节点返回，没找到时返回最后一个节点
func FindLinkNode(headNode *Cat, position int) (bool, *Cat) {
	if headNode.Next == nil {
		return false, headNode
	}
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
	return find, tmpNode
}
// 在node前面插入newNode
func AddBefore(node *Cat, newNode *Cat){
	if node.Pre != nil {
		node.Pre.Next = newNode
	}
	newNode.Pre = node.Pre
	newNode.Next = node
	node.Pre = newNode
}
// 在node后插入newNode
func AddAfter(node *Cat, newNode *Cat){
	newNode.Pre = node
	newNode.Next = node.Next
	if node.Next != nil {
		node.Next.Pre = newNode
	}
	node.Next = newNode
}
// 删除node元素
func RemoveCur(node *Cat) {
	// 一个元素循环
	if node == node.Next {
		node.Next = nil
		node.Pre = nil
	}else{
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
	}
}
// 添加结点到最后一位
func AddLastNode(headNode *Cat, newNode *Cat) {
	if headNode.Next == nil {
		headNode.Id = newNode.Id
		headNode.Next = headNode
		headNode.Pre = headNode
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
	if tmp.Next != nil {
		tmp.Next.Pre = newNode
	}
	tmp.Next = newNode
	newNode.Pre = tmp
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
	if tmp.Next != nil {
		tmp.Next.Pre = pre
	}
	pre.Next = tmp.Next
}
// position: 0 最后一位；1 第一位（是新的headNode）；没找到位置时添加到最后一位；
func AddLinkNode(headNode *Cat, newNode *Cat, position int) *Cat{
	if position == 0 {
		AddLastNode(headNode, newNode)
		return headNode
	}
	find, node := FindLinkNode(headNode, position)
	if find {
		AddBefore(node, newNode)
		if position == 1 {
			return headNode.Pre
		}
	} else {
		AddAfter(node, newNode)
	}
	return headNode
}
// position: 0 最后一位； 1第一位（headNode向下移动一位）； 没找到位置时不删除
func DelLinkNode(headNode *Cat, position int) *Cat{
	if position == 0 {
		DelLastNode(headNode)
		return headNode
	}
	find, node := FindLinkNode(headNode, position)
	if find {
		RemoveCur(node)
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