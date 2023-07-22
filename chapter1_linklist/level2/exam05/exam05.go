package exam05

type ListNode struct {
	Val int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    slow, fast := head, head
    for fast != nil && k > 1 {
        fast = fast.Next
        k--
    }
    for fast != nil && fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    return slow
}