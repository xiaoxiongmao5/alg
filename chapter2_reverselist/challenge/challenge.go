package challenge

import (
	"fmt"
	"errors"
)

type Student struct {
	Score int
	Next *Student
}

/** 思路分析
	1. 将该链表拆分到两个链表中，按照奇数位放在list1,偶数位放在list2
	2. 将list2反转
	3. 合并list1和list2，返回合并后的新链表
*/

// 链表反转
func ReverseList(head *Student) *Student{
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head
	newlist := ReverseList(tmp.Next)
	tmp.Next.Next = tmp
	tmp.Next = nil
	return newlist
}

// 拆分偶数个数链表
// list1：存放奇数位的节点	list2：存放偶数位的节点
func SplitList(head *Student) (*Student, *Student, error) {
	list1, list2 := &Student{}, &Student{}
	tmp1, tmp2 := list1, list2
	len1, len2 := 0, 0
	tmp := head
	odd := true
	for {
		next := tmp.Next
		tmp.Next = nil
		if odd {
			tmp1.Next = tmp
			tmp1 = tmp1.Next
			len1++
			odd = false
		} else {
			tmp2.Next = tmp
			tmp2 = tmp2.Next
			len2++
			odd = true
		}
		if next == nil {
			break
		}
		tmp = next
	}
	// 判断原链表非偶数个时 退出
	if len1 != len2 {
		return nil, nil, errors.New("该链表非偶数个")
	}
	return list1.Next, list2.Next, nil
}

// 合并两个链表
func MergeLists(list1 *Student, list2 *Student) *Student {
	tmp1, tmp2 := list1, list2
	odd := true
	newlist := &Student{}
	ntmp := newlist
	for {
		if tmp1 == nil && tmp2 == nil {
			break
		}
		if odd {
			next1 := tmp1.Next
			tmp1.Next = nil
			ntmp.Next = tmp1
			tmp1 = next1
			odd = false
		} else {
			next2 := tmp2.Next
			tmp2.Next = nil
			ntmp.Next = tmp2
			tmp2 = next2
			odd = true
		}
		ntmp = ntmp.Next
	}
	return newlist.Next
}

// 相互帮助
func Method(head *Student) *Student {
	if head == nil || head.Next == nil {
		return head
	}
	// 拆分偶数个数链表
	list1, list2, err := SplitList(head)
	if err != nil {
		fmt.Println("err=", err)
	}
	// 偶数链表反转
	nlist2 := ReverseList(list2)
	// 合并两个链表
	return MergeLists(list1, nlist2)
}