package exam09

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    dummyHead := &ListNode{}
    dummyHead.Next = head
    pre, tmp := dummyHead, head
    find := false
    for tmp.Next != nil {
        if tmp.Val == tmp.Next.Val {
            find = true
            tmp = tmp.Next
        } else if find == true && tmp.Val != tmp.Next.Val {
            find = false
            pre.Next = tmp.Next
            tmp = tmp.Next
        } else {
            pre = tmp
            tmp = tmp.Next
        }
    }
    if find == true {
        pre.Next = tmp.Next
    }
    return dummyHead.Next
}