package exam01

type ListNode struct {
     Val int
     Next *ListNode
}

// 使用虚拟头节点 头插法
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newlist := &ListNode{}
    // 虚拟头结点
    newlist.Next = head
    pre := newlist
    tmp := head
    // 让tmp走到left的位置,pre保持在tmp前面一个
    for i:=1; i<left; i++{
        pre = pre.Next
        tmp = tmp.Next
    }
    // 此循环 tmp会走到right的位置,但tmp节点没有变，一直是把其后一个节点拿下来放在pre后面，直到全部拿完，自己位置就处于right的位置（即：反转完区间内的所有节点了）
    for i:=0; i<right-left; i++{
        next := tmp.Next
        tmp.Next = tmp.Next.Next
        next.Next = pre.Next
        pre.Next = next
    }
    return newlist.Next
}