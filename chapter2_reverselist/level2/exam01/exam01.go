package exam01

type ListNode struct {
     Val int
     Next *ListNode
}

/** 思路分析：头插入
    1. 找到left左边的元素，保持不动（作为反转虚拟头结点），注意这里需要新创建一个节点newNode指向head，再逐步移动到left前（因为可能left是1，则虚拟头结点就是直接指向head的）
    2. 从left向后依次取下一个节点，插入到虚拟头结点后面，每次操作都相当于left结点向后走，直到占到right元素的位置
    3. 因为left元素每次去掉后面的元素，都是连接它.next.next，所以反转完之后，自然left是指向原来right元素的下一个的，所以该链表现在已经正确了
    4. return newNode.Next
*/
    
// 使用虚拟头节点 头插法
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newlist := &ListNode{}
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


/** 思路分析：穿针引线
    1. 将该链表分成3段，left前，反转区间，right后，反转完中间之后，穿针引线起来
    2. 所以需要变量结点存储（为了之后穿针引线）：left前的位置、right后的位置
    3. 在反转区间，进行一次链表反转即可（这里有三种方式：虚拟头插法、自我反转、递归反转）
*/

func reverseLink(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    tmp := head
    newlist := reverseLink(tmp.Next)
    tmp.Next.Next = tmp
    tmp.Next = nil
    return newlist
}
// 穿针引线法
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newlist := &ListNode{}
    newlist.Next = head
    pre := newlist
    // 让pre保持在left左1的位置
    for i:=1; i<left; i++{
        pre = pre.Next
    }
    tmp := pre.Next
    leftNode := tmp
    for i:=0; i<right-left; i++{
        tmp = tmp.Next
    }
    // 让rightAferNode保持在right右1的位置
    rightAferNode := tmp.Next
    // 获得反转区间
    tmp.Next = nil
    // 反转区间链表
    nleft := reverseLink(leftNode)
    leftNode.Next = rightAferNode
    pre.Next = nleft
    return newlist.Next
}

