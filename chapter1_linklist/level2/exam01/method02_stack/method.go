package method02_stack

import (
	slink "xj/chapter1_linklist/level2/exam01/single_linklist"
	stack "xj/chapter1_linklist/level2/exam01/stack"
	"fmt"
)

// 链表入栈
func List2Stack(list *slink.Cat) *stack.Stack{
	stack := &stack.Stack{
		MaxLen:10,
		Top:-1,
	}
	if list.Next == nil {
		return stack
	}
	list_tmp := list.Next
	for {
		err := stack.Push(list_tmp)	// 存储节点指针
		if err != nil {
			fmt.Println("err : ", err)
			break
		}
		if list_tmp.Next == nil {
			break
		}
		list_tmp = list_tmp.Next
	}
	return stack
}

// 使用栈
// 返回找到的公共结点，没找到时返回list1
func Method(list1 *slink.Cat, list2 *slink.Cat) (bool, *slink.Cat){
	if list1.Next == nil || list2.Next == nil {
		return false, list1
	}
	find := false
	var pre *slink.Cat
	stack1 := List2Stack(list1)
	stack2 := List2Stack(list2)
	// 出栈比较
	for {
		tmp1, err1 := stack1.Pop()
		if err1 != nil{
			fmt.Println("err1: ", err1)
			break
		}
		tmp2, err2 := stack2.Pop()
		if err2 != nil {
			fmt.Println("err2: ", err2)
			break
		}
		ntmp1, _ := tmp1.(*slink.Cat)
		ntmp2, _ := tmp2.(*slink.Cat)
		if ntmp1 != ntmp2 {
			// 结点地址相同
			find = true
			break
		} else {
			pre = ntmp1
		}
	}
	if find && pre != nil {
		return find, pre
	}
	return false, list1
}