package exam03

type ListNode struct {
	Val int
	Next *ListNode
}
 
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    newList := &ListNode{}
    tmp, tmp1, tmp2 := newList, list1, list2
    for tmp1 != nil && tmp2 != nil {
        // 两个链表都有值时才能对比大小
        if tmp1.Val >= tmp2.Val {
            tmp.Next = tmp2
            tmp2 = tmp2.Next
        } else {
            tmp.Next = tmp1
            tmp1 = tmp1.Next
        }
        tmp = tmp.Next
    }
    if tmp1 != nil {
        tmp.Next = tmp1
    }
    if tmp2 != nil {
        tmp.Next = tmp2
    }
    return newList.Next
}