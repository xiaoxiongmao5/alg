package method01_hashmap

import (
	"testing"
	"fmt"
	slink "xj/chapter1_linklist/level2/exam01/single_linklist"
)


func TestMethod(t *testing.T) {
	fmt.Printf("======= %s =======\n", "测试可以找到的数据")
	list1 := slink.BuildList(1, 2, 30, 40, 50)
	fmt.Println(slink.DumpList(list1))
	list2 := slink.BuildList(8, 9, 10, 30, 40, 50)
	fmt.Println(slink.DumpList(list2))

	find, commonNode := Method(list1, list2)
	if find && commonNode.Id == 30 {
		fmt.Printf("except\t%v \n got\t%v \n", commonNode.Id, 30)
	} else {
		t.Errorf("\n except\t%v \n got\t%v \n", commonNode.Id, 30)
	}

	fmt.Printf("======= %s =======\n", "测试找不到的数据")
	list3 := slink.BuildList(1, 2, 300, 400, 500)
	fmt.Println(slink.DumpList(list3))
	list4 := slink.BuildList(8, 9, 10, 30, 40, 50)
	fmt.Println(slink.DumpList(list4))
	
	find, commonNode = Method(list3, list4)
	if !find {
		fmt.Printf("except \t%v \n got\t%v \n", false, find)
	} else {
		t.Errorf("\n except \t%v \n got\t%v \n", false, find)
	}

	fmt.Printf("======= %s =======\n", "测试传递空链表")
	list5 := slink.BuildList()
	fmt.Println(slink.DumpList(list5))
	list6 := slink.BuildList(8, 9, 10, 30, 40, 50)
	fmt.Println(slink.DumpList(list6))
	
	find, commonNode = Method(list3, list4)
	if !find {
		fmt.Printf("except \t%v \n got\t%v \n", false, find)
	} else {
		t.Errorf("\n except \t%v \n got\t%v \n", false, find)
	}
}