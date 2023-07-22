package exam04

type ListNode struct {
	Val int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    slow := head
    fast := head
    for {
        if fast == nil || fast.Next == nil {
            break   
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}