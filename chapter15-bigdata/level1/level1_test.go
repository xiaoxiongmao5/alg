package level1

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

func TestFindDuplicateByByteIn32000(t *testing.T) {
	arr := []int{1, 2, 3, 3, 4, 5, 6, 6, 7, 7, 8, 9, 10, 15, 15, 30, 30, 45, 99, 99, 100}
	res := FindDuplicateByByteIn32000(arr)
	equal[[]int](t, "FindDuplicateByByteIn32000", res, []int{3, 6, 7, 15, 30, 99})
}

func TestFindDuplicateByIntIn32000(t *testing.T) {
	arr := []int{1, 2, 3, 3, 4, 5, 6, 6, 7, 7, 8, 9, 10, 15, 15, 30, 30, 45, 99, 99, 100}
	res := FindDuplicateByIntIn32000(arr)
	equal[[]int](t, "FindDuplicateByIntIn32000", res, []int{3, 6, 7, 15, 30, 99})
}
