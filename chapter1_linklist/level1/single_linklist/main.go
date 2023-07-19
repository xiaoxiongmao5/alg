package single_linklist

import (
	"fmt"
	"strings"
)

// type LinkNode[T any] struct {
// 	Next *T
// }

// type LinkList[T any] struct {
// 	Head *LinkNode[T]
	
// }

type Cat struct {
	Id int
	Next *Cat
}

// 找到position前一个节点，没找到时返回最后一个节点
func FindLinkNode(headNode *Cat, position int) (bool, *Cat){
	if headNode.Next == nil {
		return false, headNode
	}
	tmpNode := headNode
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
func AddAfter(node *Cat, newNode *Cat) {
	newNode.Next = node.Next
	node.Next = newNode
}
// 删除node的下一个元素
func DelNodeAfter(node *Cat) {
	if node.Next != nil {
		node.Next = node.Next.Next
	}
}
// 添加结点到最后一位
func AddLastNode(headNode *Cat, newNode *Cat) {
	if headNode == nil {
		headNode.Next = newNode
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
}
// 删除最后一个结点
func DelLastNode(headNode *Cat) {
	if headNode.Next == nil {
		return
	}
	tmp, pre := headNode, headNode
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
func AddLinkNode(headNode *Cat, newNode *Cat, position int) {
	if position == 0 {
		AddLastNode(headNode, newNode)
		return
	}
	_, tmpNode := FindLinkNode(headNode, position)
	AddAfter(tmpNode, newNode)
}
// position: 0 最后一位； 1第一位； 没找到位置时不删除
func DelLinkNode(headNode *Cat, position int){
	if position == 0 {
		DelLastNode(headNode)
		return
	}
	find, tmpNode := FindLinkNode(headNode, position)
	if find {
		DelNodeAfter(tmpNode)
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

func BuildList(values ...int) *Cat{
	head := &Cat{}
	tmp := head

	for _, v := range values {
		node := &Cat{Id: v}
		tmp.Next = node
		tmp = node
	}
	return head
}

func DumpList(headNode *Cat) string{
	if headNode.Next == nil {
		return "empty"
	}
	var ret = make([]string, 0)
	tmp := headNode.Next
	id := 1
	for {
		ret = append(ret, fmt.Sprintf("%v[%v]", id, tmp.Id))
		tmp = tmp.Next
		id++
		if tmp == nil {
			break
		}
	}
	return strings.Join(ret, "=>")
}