package reverselist

import (
	"fmt"
	"testing"
	slink "xj/chapter2_reverselist/level1/single_linklist"
)

func equal[T any](t *testing.T, desc string, list1 *slink.LinkNode[T], list2 *slink.LinkNode[T]) {
	str1 := slink.DumpListNoHead(list1)
	str2 := slink.DumpListNoHead(list2)

	if str1 == str2 {
		fmt.Printf("======= %s 通过 =======\n", desc)
		fmt.Printf("except:\t%s \n got:\t%s \n", str2, str1)
	} else {
		fmt.Printf("======= %s 未通过 =======\n", desc)
		t.Errorf("\n except:\t%s \n got:\t%s \n", str2, str1)
	}
}

// 使用辅助头结点反转
func TestReverselistByHead(t *testing.T) {
	fmt.Printf("======= %s =======\n", "【【 测试：使用头结点反转 】】")
	list := slink.BuildList(1,2,3,4,5)
	equal(t, "无重复数据", ReverselistByHead(list), slink.BuildList(5,4,3,2,1))

	list2 := slink.BuildList(1,2,2,4,5)
	equal(t, "有重复值", ReverselistByHead(list2), slink.BuildList(5,4,2,2,1))

	list3 := slink.BuildList(5)
	equal(t, "一个数据", ReverselistByHead(list3), slink.BuildList(5))

	list4 := slink.BuildList[int]()
	equal(t, "无数据", ReverselistByHead(list4), slink.BuildList[int]())
}

// 自身反转
func TestReverselistBySelf(t *testing.T) {
	fmt.Printf("======= %s =======\n", "【【 测试：自身反转 】】")
	list := slink.BuildListNoHead(1,2,3,4,5)
	equal(t, "无重复数据", ReverselistBySelf(list), slink.BuildListNoHead(5,4,3,2,1))

	list2 := slink.BuildListNoHead(1,2,2,4,5)
	equal(t, "有重复值", ReverselistBySelf(list2), slink.BuildListNoHead(5,4,2,2,1))

	list3 := slink.BuildListNoHead(5)
	equal(t, "一个数据", ReverselistBySelf(list3), slink.BuildListNoHead(5))

	list4 := slink.BuildListNoHead[int]()
	equal(t, "无数据", ReverselistBySelf(list4), slink.BuildListNoHead[int]())
}

// 使用递归反转
func TestReverseListByRecursion(t *testing.T) {
	fmt.Printf("======= %s =======\n", "【【 测试：使用递归反转 】】")
	list := slink.BuildListNoHead(1,2,3,4,5)
	equal(t, "无重复数据", ReverseListByRecursion(list), slink.BuildListNoHead(5,4,3,2,1))

	list2 := slink.BuildListNoHead(1,2,2,4,5)
	equal(t, "有重复值", ReverseListByRecursion(list2), slink.BuildListNoHead(5,4,2,2,1))

	list3 := slink.BuildListNoHead(5)
	equal(t, "一个数据", ReverseListByRecursion(list3), slink.BuildListNoHead(5))

	list4 := slink.BuildListNoHead[int]()
	equal(t, "无数据", ReverseListByRecursion(list4), slink.BuildListNoHead[int]())
}