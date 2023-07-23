package exam04

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head
	newlist := reverseList(tmp.Next)
	tmp.Next.Next = tmp
	tmp = nil
	return newlist
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummyHead := &ListNode{}
	tmp := dummyHead
	l1, l2 = reverseList(l1), reverseList(l2)
	site_shi, val, val1, val2 := 0, 0, 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		val = val1 + val2 + site_shi
		if val >= 10 {
			val = val % 10
			site_shi = 1
		} else {
            site_shi = 0
        }
		tmp.Next = &ListNode{Val:val}
		tmp = tmp.Next
		val1, val2 = 0, 0
	}
	if site_shi == 1 {
		tmp.Next = &ListNode{Val:site_shi}
	}
	return reverseList(dummyHead.Next)
}