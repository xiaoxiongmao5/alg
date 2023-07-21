package exam02

/** 解法思想：
	1. 让虚拟头结点指向链表，设置前置节点pre指向第一个节点
	2. 从第一个节点开始作为当前节点tmp，指向其.next.next节点，目的是去掉其后的节点（注意，要先提前保存好这个要去掉的结点），然后将去掉的节点插入到当前节点前(即：pre节点后)，这就完成了当前节点的反转。
	3. 然后让pre=tmp，tmp=tmp.Next(即：当前节点向下走一个），接着参照步骤2，与其后的节点进行反转
	4. 注意：停止循环的条件，要考虑到tmp.Next.Next是否为nil，因为只有tmp.Next.Next存在的时候，才需要把tmp.Next和tmp自己进行交换
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dummyNode := &ListNode{}
    dummyNode.Next = head
    pre := dummyNode
    tmp := head
    for {
        next := tmp.Next
        tmp.Next = tmp.Next.Next
        next.Next = tmp
        pre.Next = next
        // 注意tmp.Next.Next == nil条件，因为两两交换tmp需要指向tmp.Next.Next才能把tmp后面的值取下来进行交换
        if tmp.Next == nil || tmp.Next.Next == nil{
            break
        }
        pre = tmp
        tmp = tmp.Next
    }
    return dummyNode.Next
}