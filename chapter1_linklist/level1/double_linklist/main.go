package double_linklist

import (
	"fmt"
)

type Cat struct {
	Id int
	Pre *Cat
	Next *Cat
}

// 找到position的当前节点返回，没找到时返回最后一个节点
func FindLinkNode(headNode *Cat, position int) (bool, *Cat){
	if headNode.Next == nil {
		return false, headNode
	}
	tmpNode := headNode
	len := 0
	find := false
	for {
		len++
		tmpNode = tmpNode.Next
		if len == position {
			find = true
			break
		}
		if tmpNode.Next == nil {
			break
		}
	}
	return find, tmpNode
}
// 在node前插入newNode
func AddBefore(node *Cat, newNode *Cat) {
	if node.Pre != nil {
		node.Pre.Next = newNode
	}
	newNode.Next = node
	newNode.Pre = node.Pre
	node.Pre = newNode
}
// 在node后插入newNode
func AddAfter(node *Cat, newNode *Cat) {
	newNode.Next = node.Next
	if node.Next != nil {
		node.Next.Pre = newNode
	}
	node.Next = newNode
	newNode.Pre = node
}
// 删除node节点
func DelNode(node *Cat) {
	if node.Pre != nil {
		node.Pre.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Pre = node.Pre
	}
}
// 添加结点到最后一位
func AddLastNode(headNode *Cat, newNode *Cat) {
	if headNode == nil {
		headNode.Next = newNode
		newNode.Pre = headNode
		return
	}
	tmp := headNode
	for {
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	tmp.Next = newNode
	newNode.Pre = tmp
}
// 删除最后一个结点
func DelLastNode(headNode *Cat) {
	if headNode.Next == nil {
		return
	}
	tmp := headNode
	for {
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	if tmp.Pre != nil {
		tmp.Pre.Next = nil
	}
}
// position: 0 最后一位；1 第一位；没找到位置时添加到最后一位；
func AddLinkNode(headNode *Cat, newNode *Cat, position int) {
	if position == 0 {
		AddLastNode(headNode, newNode)
		return
	}
	find, tmpNode := FindLinkNode(headNode, position)
	if find  {
		AddBefore(tmpNode, newNode)
	} else {
		AddAfter(tmpNode, newNode)
	}
}
// position: 0 最后一位； 1第一位； 没找到位置时不删除
func DelLinkNode(headNode *Cat, position int){
	if position == 0 {
		DelLastNode(headNode)
		return
	}
	find, tmpNode := FindLinkNode(headNode, position)
	if find {
		DelNode(tmpNode)
	}
}

func ShowLink(headNode *Cat) {
	if headNode.Next == nil {
		fmt.Println("link empty")
		fmt.Println()
		return
	}
	tmpNode := headNode
	for {
		fmt.Printf("Id=%d ==>\n", tmpNode.Next.Id)
		tmpNode = tmpNode.Next
		if tmpNode.Next == nil {
			break
		}
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
		tmpNode = tmpNode.Next
		if tmpNode.Next == nil {
			break
		}
	}
	for {
		fmt.Printf("Id=%d ==>\n", tmpNode.Id)
		tmpNode = tmpNode.Pre
		if tmpNode.Pre == nil {
			break
		}
	}
	fmt.Println()
}