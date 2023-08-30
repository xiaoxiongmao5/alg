package challenge

import (
	"fmt"
	"reflect"
	"testing"
)

func equal[T string | int | []int](t *testing.T, desc string, str1 T, str2 T) {
	fmt.Printf("======= %s =======\n", desc)
	if reflect.DeepEqual(str1, str2) {
		fmt.Printf("\n通过\n export\t%v\n got\t%v\n", str2, str1)
	} else {
		t.Errorf("\n未通过\n export\t%v\n got\t%v\n", str2, str1)
	}
	fmt.Println("")
}

func TestFindNoExistNumsBy1G(t *testing.T) {
	arr := []int{0, 1, 5, 7, 8, 9, 19}
	res := FindNoExistNumsBy1G(arr)
	equal[int](t, "FindNoExistNumsBy1G", len(res), 4000000000-7)
}
