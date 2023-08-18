package challenge

import (
	"fmt"
	"strings"
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

func TestCompress(t *testing.T) {
	arr := []string{"tiantian", "yuanyuan", "tingting", "yingzi", "qiyue"}
	compareArr := []string{"tian~", "yuan~", "ting~", "yingzi", "qiyue"}
	Compress(arr)
	equal[string](t, "名字缩写", strings.Join(arr, ","), strings.Join(compareArr, ","))
}
