package challenge

import (
	"fmt"
	"testing"
)

func equal[T string | int](t *testing.T, desc string, str1 T, str2 T) {
	fmt.Printf("======= %s =======\n", desc)
	if str1 == str2 {
		fmt.Printf("\n通过\n export\t%v\n got\t%v\n", str2, str1)
	} else {
		t.Errorf("\n未通过\n export\t%v\n got\t%v\n", str2, str1)
	}
	fmt.Println("")
}

func TestMethod(t *testing.T) {
	equal[int](t, "测试5", Method(5), 3)
	equal[int](t, "测试15", Method(15), 13)
	equal[int](t, "测试25", Method(25), 23)
	equal[int](t, "测试50", Method(50), 39)
	equal[int](t, "测试500", Method(500), 399)
	equal[int](t, "测试5000", Method(5000), 3999)
}
