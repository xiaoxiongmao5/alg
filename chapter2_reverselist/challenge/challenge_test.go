package challenge

import (
	"testing"
	"fmt"
	"strings"
)

func initListData(num int) *Student{
	dummyNode := &Student{}
	tmp := dummyNode
	for i:=1; i<=num; i++ {
		node := &Student{Score:i}
		tmp.Next = node
		tmp = tmp.Next
	}
	return dummyNode.Next
}

func exportListData(values ...int) *Student{
	dummyNode := &Student{}
	tmp := dummyNode
	for _, v := range values {
		node := &Student{Score:v}
		tmp.Next = node
		tmp = tmp.Next
	}
	return dummyNode.Next
}

func dumpList(list *Student) string {
	var ret = make([]string, 0)
	tmp := list
	for {
		if tmp == nil {
			break
		}
		ret = append(ret, fmt.Sprintf("%v", tmp.Score))
		tmp = tmp.Next
	}
	return strings.Join(ret, "=>")
}

func equal(t *testing.T, desc string, str1 string, str2 string) {
	fmt.Printf("======= %s =======\n", desc)
	if str1 == str2 {
		fmt.Printf("\n通过\n export\t%v\n got\t%v\n", str2, str1)
	} else {
		t.Errorf("\n未通过\n export\t%v\n got\t%v\n", str2, str1)
	}
}

func TestMethod(t *testing.T){
	stuList := initListData(10)
	newStuList := Method(stuList)
	exportList := exportListData(1, 10, 3, 8, 5, 6, 7, 4, 9, 2)
	equal(t, "测试相互帮助", dumpList(newStuList), dumpList(exportList))
}