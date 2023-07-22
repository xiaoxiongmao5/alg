package exam07

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return head
    }
    dummyHead := &ListNode{}
    dummyHead.Next = head
    tmp := dummyHead
    for {
        if tmp.Next == nil {
            break
        }
        if tmp.Next.Val == val {
            tmp.Next = tmp.Next.Next
        } else {
            tmp = tmp.Next
        }
    }
    return dummyHead.Next
}