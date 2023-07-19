package method04_sub

import (
	slink "xj/chapter1_linklist/level2/exam01/single_linklist"
	// "fmt"
)

// 使用消除步差
// 返回找到的公共结点，没找到时返回list1
func Method(list1 *slink.Cat, list2 *slink.Cat) (bool, *slink.Cat){
	if list1.Next == nil || list2.Next == nil {
		return false, list1
	}
	list1_tmp := list1.Next
	list2_tmp := list2.Next
	find := false

	len1 := slink.GetListLen(list1)
	len2 := slink.GetListLen(list2)
	sub := 0
	// 消除步差
	if len1 >= len2 {
		sub = len1 - len2
		for ; sub>0; sub-- {
			list1_tmp = list1_tmp.Next
		}
	} else {
		sub = len2 - len1
		for ; sub>0; sub-- {
			list2_tmp = list2_tmp.Next
		}
	}
	for {
		// 找到了
		if list1_tmp == list2_tmp {
			find = true
			break
		}
		if list1_tmp.Next == nil {
			break
		}
		list1_tmp = list1_tmp.Next
		list2_tmp = list2_tmp.Next
	}
	if find {
		return find, list2_tmp
	}
	return find, list1
}