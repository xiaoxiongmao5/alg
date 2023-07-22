package exam02

/** 思路分析
	1. 使用快慢指针，当快指针到达结尾停下时，慢指针正好在中间，同时慢指针走的每一步，都将当前节点存储起来，确保逆序存储（比如：栈、数组、链表）
	2. 接着比较从存储区域依次取值和慢指针往下走取值，如果有不等的，就不是回文序列，否则是回文序列。（注意：这里比较的是值内容，而不是该节点指针地址）
*/

/** 注意事项
1. 使用快慢指针的时候，注意链表长度是单数双数的问题。
	* 如果是单数，当fast.Next == nil停下时，slow在中间位置，包含中间数
	* 如果是双数，放fast == nil停下时，slow在中间位置的后一个位置
	1		2		3		nil
	slow	slow(当因fast终止时，这里的slow对应操作还没完成，需要补上这一步)
	fast			fast	
	1		2		3		4	nil
	slow	slow	slow		
	fast			fast		fast	
*/

type ListNode struct {
	Val int
	Next *ListNode
}
// 在头部添加新节点，并返回新头部节点
func addHeadNode(head *ListNode, newNode *ListNode) *ListNode {
    node := &ListNode{
        Val:newNode.Val,
    }
    if head == nil {
        head = node
        return head
    }
    node.Next = head
    head = node
    return head
}
func isPalindrome(head *ListNode) bool {
    if head == nil {
        return false
    }
    if head.Next == nil {
        return true
    }
    slow := head
    fast := head
    var newHead *ListNode
    for {
        if fast == nil {
            break
        }
        if fast.Next == nil {
            // 注意这里：当链表为奇数个数时，fast正好是最后一个，slow处于正中间，此时要记得把slow放在存储链表中
            newHead = addHeadNode(newHead, slow)
            break
        }
        newHead = addHeadNode(newHead, slow)
        slow = slow.Next
        fast = fast.Next.Next
    }
    isSame := true
    for {
        if slow == nil {
            break
        }
        if slow.Val != newHead.Val {
            isSame = false
            break
        }
        slow = slow.Next
        newHead = newHead.Next
    }
    return isSame
}
// ————————————————————————————————————————————————————————————————
/** 解法二:思路（快慢指针+递归反转链表）
	1. 使用快慢指针，当快指针到达结尾停下时，慢指针正好在中间（奇数链表时，慢指针在正中间；偶数链表时，慢指针在中间的下一位）
	2. 然后反转此时的慢指针链表（比如：1->2->3->2->1，反转3->2->1，得到1(注意这个1是原链表的尾节点指针)->2->3）
	3. 最后循环反转后的链表和原链表，依次取出首节点对比，值不等说明不是回文序列，直达循环结束，结束条件是反转链表遍历完，如果到结束都是相等的，说明就是回文序列

	时间复杂度: O(n)
	空间复杂度: O(1)
*/
func reverselist(head *ListNode) *ListNode {
    fmt.Println("reverse:", head.Val)
    if head == nil || head.Next == nil {
        return head
    }
    tmp := head
    newHead := reverselist(tmp.Next)
    tmp.Next.Next = tmp
    tmp.Next = nil
    return newHead
}
func isPalindrome2(head *ListNode) bool {
    if head == nil {
        return false
    }
    if head.Next == nil {
        return true
    }
    slow := head
    fast := head
    var newList *ListNode
    for {
        if fast == nil || fast.Next == nil{
            break
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    newList = reverselist(slow)
    isSame := true
    for {
        if newList == nil || head == nil {
            break
        }
        if newList.Val != head.Val {
            isSame = false
            break
        }
        newList = newList.Next
        head = head.Next
    }
    return isSame
}