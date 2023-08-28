package level2

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

func TestFindNoExistNumBy1G(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13}
	res := FindNoExistNumBy1G(arr)
	equal[int](t, "FindNoExistNumBy1G", res, 8)
}

func TestFindNoExistNumBy10M(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13}
	res := FindNoExistNumBy10M(arr)
	equal[int](t, "FindNoExistNumBy10M", res, 12)
}
