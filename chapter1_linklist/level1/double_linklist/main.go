package double_linklist

import (
	"fmt"
)

type Cat struct {
	Id int
	Pre *Cat
	Next *Cat
}

// position: 0 最后一位；1 第一位；没找到位置时添加到最后一位；
func AddLinkNode(headNode *Cat, newNode *Cat, position int) {
	// 若该链表是否为空
	if headNode.Next == nil {
		headNode.Next = newNode
		newNode.Pre = headNode
		return
	}
	tmpNode := headNode
	// 若添加位置为1，插入在首位（headNode后第一个）
	if position == 1 {
		newNode.Next = tmpNode.Next
		tmpNode.Next.Pre = newNode
		tmpNode.Next = newNode
		newNode.Pre = tmpNode
		return
	}

	len := 1
	for {
		if tmpNode.Next == nil {
			break
		}
		tmpNode = tmpNode.Next
		len++
		// 找到添加位置了
		if len == position {
			break
		}
	}
	newNode.Next = tmpNode.Next
	if tmpNode.Next != nil {
		tmpNode.Next.Pre = newNode
	}
	tmpNode.Next = newNode
	newNode.Pre = tmpNode
}

// position: 0 最后一位； 1第一位； 没找到位置时不删除
func DelLinkNode(headNode *Cat, position int){
	if headNode.Next == nil {
		return
	}
	tmpNode := headNode
	len := 1
	flag := false
	for {
		if position == len {
			flag = true
			break
		}
		if tmpNode.Next.Next == nil {
			break
		}
		len ++
		tmpNode = tmpNode.Next
		
	}
	if flag {
		if tmpNode.Next.Next != nil {
			tmpNode.Next.Next.Pre = tmpNode
		}
		tmpNode.Next = tmpNode.Next.Next
		return
	}
	if position == 0 {
		tmpNode.Next = nil
		return
	}
	fmt.Println("position越界")
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