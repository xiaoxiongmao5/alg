package single_linklist

import "testing"
import "fmt"

func equal(t *testing.T, desc string, l1 *LinkNode[int], l2 *LinkNode[int]) {
	var l1_dump = DumpList(l1)
	var l2_dump = DumpList(l2)
	var result = l1_dump == l2_dump
	var result_str = "成功"
	if !result {
		result_str = "失败!!!"
	}
	fmt.Printf("======= %s[%s] =======\n", desc, result_str)
	if result {
		fmt.Printf("except\t%v \n got\t%v \n", l2_dump, l1_dump)
	} else {
		t.Errorf(" %s \n except\t%v \n got\t%v \n", desc, l2_dump, l1_dump)
	}
}

func TestAddLinkNode(t *testing.T) {
	fmt.Println("======= 开始测试包: single_linklist =======")
	head := &LinkNode[int]{}
	list := &LinkList[int]{
		Head : head,
	}

	cat1 := &LinkNode[int]{Val: 1}
	list.AddLinkNode(cat1, 0)
	equal(t, "添加到空链表的末尾", head, BuildList(1))

	cat2 := &LinkNode[int]{Val: 2}
	list.AddLinkNode(cat2, 5)
	equal(t, "添加到不存在的位置5", head, BuildList(1, 2))

	cat3 := &LinkNode[int]{Val: 3}
	list.AddLinkNode(cat3, 1)
	equal(t, "添加到头部", head, BuildList(3, 1, 2))

	cat4 := &LinkNode[int]{Val: 4}
	list.AddLinkNode(cat4, 2)
	equal(t, "添加到中间第二个节点", head, BuildList(3, 4, 1, 2))

	cat5 := &LinkNode[int]{Val: 5}
	list.AddLinkNode(cat5, 15)
	equal(t, "添加到不存在位置15", head, BuildList(3, 4, 1, 2, 5))

	list.DelLinkNode(0)
	equal(t, "删除0", head, BuildList(3, 4, 1, 2))

	list.DelLinkNode(1)
	equal(t, "删除1", head, BuildList(4, 1, 2))

	list.DelLinkNode(100)
	equal(t, "删除100", head, BuildList(4, 1, 2))

	list.DelLinkNode(2)
	equal(t, "删除2", head, BuildList(4, 2))

	list.DelLinkNode(2)
	equal(t, "删除2", head, BuildList(4))

	list.DelLinkNode(0)
	equal(t, "删除0", head, BuildList[int]())

	list.DelLinkNode(1)
	equal(t, "删除1", head, BuildList[int]())
}
