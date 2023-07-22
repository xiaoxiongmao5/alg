package exam08

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {
        return head
    }
    dummyHead := &ListNode{}
    dummyHead.Next = head
    slow, fast := dummyHead, dummyHead
    // 走到正数n后面的位置
    // fast先走n步，当fast走到队尾的时候，slow就是倒数n节点的前一个节点
    for i:=1; i<n+1; i++ {
        fast = fast.Next
        if fast == nil {
            break
        }
    }
    if fast == nil {
        return head
    }

    for fast.Next != nil {
        slow = slow.Next
        fast = fast.Next
    }
    slow.Next = slow.Next.Next
    return dummyHead.Next
}