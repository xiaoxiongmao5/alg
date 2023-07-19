package method01_hashmap

import (
	slink "xj/chapter1_linklist/level2/exam01/single_linklist"
)

// 使用hashMap
// 返回找到的公共结点，没找到时返回list1
func Method(list1 *slink.Cat, list2 *slink.Cat) (bool, *slink.Cat){
	if list1.Next == nil || list2.Next == nil {
		return false, list1
	}
	list1_map := make(map[int]*slink.Cat)
	list1_tmp := list1.Next
	// 遍历list1，放到hashmap
	for {
		list1_map[list1_tmp.Id] = list1_tmp
		if list1_tmp.Next == nil {
			break
		}
		list1_tmp = list1_tmp.Next
	}
	find := false
	list2_tmp := list2.Next
	// 遍历list2，找到hashmap中存在的点 就是公共结点
	for {
		_, ok := list1_map[list2_tmp.Id]
		if ok {
			find = true
			break
		} else {
			list2_tmp = list2_tmp.Next
		}
		if list2_tmp == nil {
			break
		}
	}
	if find {
		// fmt.Printf("找到的公共点的Id为 %d \n", list2_tmp.Id)
		return find, list2_tmp
	}
	// fmt.Println("这两个链表没有公共的点")
	return find, list1
}