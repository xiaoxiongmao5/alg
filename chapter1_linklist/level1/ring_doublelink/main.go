package ring_doublelink

import (
	"fmt"
)

type Cat struct {
	Id int
	Pre *Cat
	Next *Cat
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
	tmpNode := headNode
	lastNode := headNode
	for {
		lastNode = lastNode.Next
		if lastNode.Next == headNode {
			break
		}
	}
	// 若添加位置为1，插入在首位（作为新的headNode）
	if position == 1 {
		lastNode.Next = newNode
		newNode.Next = tmpNode
		newNode.Pre = lastNode
		tmpNode.Pre = newNode
		return newNode
	}

	len := 1
	for {
		tmpNode = tmpNode.Next
		lastNode = lastNode.Next
		len++
		// 找到添加位置了
		if len == position {
			break
		}
		if tmpNode == headNode {
			break
		}
	}
	newNode.Next = tmpNode
	lastNode.Next = newNode
	newNode.Pre = lastNode
	tmpNode.Pre = newNode
	return headNode
}

// position: 0 最后一位； 1第一位（headNode向下移动一位）； 没找到位置时不删除
func DelLinkNode(headNode *Cat, position int) *Cat{
	if headNode.Next == nil {
		return headNode
	}
	tmpNode := headNode
	lastNode := headNode
	for {
		lastNode = lastNode.Next
		if lastNode.Next == headNode {
			break
		}
	}
	// 说明当前只有一个结点
	if lastNode == tmpNode {
		if position == 0 || position == 1 {
			headNode.Next = nil
			headNode.Pre = nil
			return headNode
		}
	}
	// 删除第一位，headNode要向下移动一位
	if position == 1 {
		lastNode.Next = tmpNode.Next
		tmpNode.Next.Pre = lastNode
		tmpNode = tmpNode.Next
		return tmpNode
	}

	len := 1
	flag := false
	for {
		tmpNode = tmpNode.Next
		lastNode = lastNode.Next
		len++
		if position == len {
			flag = true
			break
		}
		if tmpNode.Next == headNode {
			break
		}
	}
	// 找到位置 或者 删除的是最后一位
	if flag || position == 0{
		lastNode.Next = tmpNode.Next
		tmpNode.Next.Pre = lastNode
		return headNode
	}
	fmt.Println("position越界")
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