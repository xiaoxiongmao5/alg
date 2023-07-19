package double_linklist

import "testing"
import "fmt"

func equal(t *testing.T, desc string, l1 *Cat, l2 *Cat) {
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
	head := &Cat{}

	cat1 := &Cat{Id: 1}
	AddLinkNode(head, cat1, 0)
	equal(t, "添加到空链表的末尾", head, BuildList(1))

	cat2 := &Cat{Id: 2}
	AddLinkNode(head, cat2, 5)
	equal(t, "添加到不存在的位置5", head, BuildList(1, 2))

	cat3 := &Cat{Id: 3}
	AddLinkNode(head, cat3, 1)
	equal(t, "添加到头部", head, BuildList(3, 1, 2))

	cat4 := &Cat{Id: 4}
	AddLinkNode(head, cat4, 2)
	equal(t, "添加到中间第二个节点", head, BuildList(3, 4, 1, 2))

	cat5 := &Cat{Id: 5}
	AddLinkNode(head, cat5, 15)
	equal(t, "添加到不存在位置15", head, BuildList(3, 4, 1, 2, 5))

	DelLinkNode(head, 0)
	equal(t, "删除0", head, BuildList(3, 4, 1, 2))

	DelLinkNode(head, 1)
	equal(t, "删除1", head, BuildList(4, 1, 2))

	DelLinkNode(head, 100)
	equal(t, "删除100", head, BuildList(4, 1, 2))

	DelLinkNode(head, 2)
	equal(t, "删除2", head, BuildList(4, 2))

	DelLinkNode(head, 2)
	equal(t, "删除2", head, BuildList(4))

	DelLinkNode(head, 0)
	equal(t, "删除0", head, BuildList())

	DelLinkNode(head, 1)
	equal(t, "删除1", head, BuildList())
}
