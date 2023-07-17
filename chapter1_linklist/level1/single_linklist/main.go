package single_linklist

import (
	"fmt"
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

func FindLinkNode(headNode *Cat, position int) *Cat{
	tmpNode := headNode
	pre := headNode
	len := 0
	for {
		if tmpNode.Next == nil {
			break
		}
		// 找到添加位置了
		if len + 1 == position {
			break
		}
		len++
		pre = tmpNode
		tmpNode = tmpNode.Next
	}
	if position == 0 {
		return pre
	}
	return tmpNode
}

// position: 0 最后一位；1 第一位；没找到位置时添加到最后一位；
func AddLinkNode(headNode *Cat, newNode *Cat, position int) {

	tmpNode := FindLinkNode(headNode, position)
	newNode.Next = tmpNode.Next
	tmpNode.Next = newNode
}

// position: 0 最后一位； 1第一位； 没找到位置时不删除
func DelLinkNode(headNode *Cat, position int){

	tmpNode := FindLinkNode(headNode, position)
	nextNode := tmpNode.Next
	if nextNode != nil {
		tmpNode.Next = nextNode.Next
	}else{
		tmpNode.Next = nil
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