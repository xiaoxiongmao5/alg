package challenge

import (
	"testing"
	"fmt"
)

func equal(t *testing.T, desc string, str1 string, str2 string) {
	fmt.Printf("======= %s =======\n", desc)
	if str1 == str2 {
		fmt.Printf("\n通过\n export:\t%s\n got:\t%s\n", str2, str1)
	} else {
		t.Errorf("\n未通过\n export:\t%s\n got:\t%s\n", str2, str1)
	}
}

func arr2str(arr []int, length int) string {
	arr = arr[:length]
	return fmt.Sprintf("%v", arr)
}

func TestMethod(t *testing.T) {
	arr1 := []int{1,1,2,2,3,4,5}
	length1 := Method(arr1)
	equal(t, "去掉重复元素", arr2str(arr1,length1), arr2str([]int{3,4,5}, 3))

	arr2 := []int{0,2,2,2,2,2,5}
	length2 := Method(arr2)
	equal(t, "去掉重复元素", arr2str(arr2,length2), arr2str([]int{0,5}, 2))
}