package method03_join

import (
	"testing"
	"fmt"
	slink "xj/chapter1_linklist/level2/exam01/single_linklist"
)


func TestMethod(t *testing.T) {
	fmt.Printf("======= %s =======\n", "测试可以找到的数据")
	node1 := &slink.Cat{Id: 1}
	node2 := &slink.Cat{Id: 2}
	node3 := &slink.Cat{Id: 30}
	node4 := &slink.Cat{Id: 40}
	node5 := &slink.Cat{Id: 50}
	list1 := &slink.Cat{}
	list1.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node8 := &slink.Cat{Id: 8}
	node9 := &slink.Cat{Id: 9}
	node10 := &slink.Cat{Id: 10}
	list2 := &slink.Cat{}
	list2.Next = node8
	node8.Next = node9
	node9.Next = node10
	node10.Next = node3
	node3.Next = node4
	node4.Next = node5

	fmt.Println(slink.DumpList(list1))
	fmt.Println(slink.DumpList(list2))

	find, commonNode := Method(list1, list2)
	if find && commonNode.Id == 30 {
		fmt.Printf("except\t%v \n got\t%v \n", 30, commonNode.Id)
	} else {
		t.Errorf("\n except\t%v \n got\t%v \n", 30, commonNode)
	}

	fmt.Printf("======= %s =======\n", "测试找不到的数据")
	list3 := slink.BuildList(1, 2, 30, 40, 50)
	fmt.Println(slink.DumpList(list3))
	list4 := slink.BuildList(8, 9, 10, 30, 40, 50)
	fmt.Println(slink.DumpList(list4))
	
	find, commonNode = Method(list3, list4)
	if !find {
		fmt.Printf("except\t%v \n got\t%v \n", false, find)
	} else {
		t.Errorf("\n except\t%v \n got\t%v \n", false, find)
	}

	fmt.Printf("======= %s =======\n", "测试传递空链表")
	list5 := slink.BuildList()
	fmt.Println(slink.DumpList(list5))
	list6 := slink.BuildList(8, 9, 10, 30, 40, 50)
	fmt.Println(slink.DumpList(list6))
	
	find, commonNode = Method(list5, list6)
	if !find {
		fmt.Printf("except\t%v \n got\t%v \n", false, find)
	} else {
		t.Errorf("\n except\t%v \n got\t%v \n", false, find)
	}
}