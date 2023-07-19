package method03_join

import (
	slink "xj/chapter1_linklist/level2/exam01/single_linklist"
	"fmt"
)

// 使用拼接对方
// 返回找到的公共结点，没找到时返回list1
func Method(list1 *slink.Cat, list2 *slink.Cat) (bool, *slink.Cat){
	if list1.Next == nil || list2.Next == nil {
		return false, list1
	}
	list1_tmp := list1.Next
	list2_tmp := list2.Next
	find := false
	list1_joined := false
	list2_joined := false
	for {
		// 找到了
		if list1_tmp == list2_tmp {
			find = true
			break
		}
		list1_tmp = list1_tmp.Next
		list2_tmp = list2_tmp.Next
		
		if list1_tmp == nil && list2_tmp == nil {
			break
		}

		// 拼接对方
		if list1_tmp == nil && !list1_joined{
			list1_joined = true
			list1_tmp = list2.Next
		}
		if list2_tmp == nil && !list2_joined{
			list2_joined = true
			list2_tmp = list1.Next
		}
	}
	if find {
		fmt.Printf("找到的公共点的Id为 %d \n", list2_tmp.Id)
		return find, list2_tmp
	}
	fmt.Println("这两个链表没有公共的点")
	return find, list1
}