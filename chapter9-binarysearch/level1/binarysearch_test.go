package level1

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

func TestSearch(t *testing.T) {
	arr := []int{-1, 0, 3, 5, 9, 12}
	equal[int](t, "测试数组中目标,递归,存在", SearchByRecursion(arr, 9), 4)
	equal[int](t, "测试数组中目标,递归,不存在", SearchByRecursion(arr, 2), -1)

	equal[int](t, "测试数组中目标,循环,存在", SearchByFor(arr, 9), 4)
	equal[int](t, "测试数组中目标,循环,不存在", SearchByFor(arr, 2), -1)
}

func TestSearchLeftInRepeatArr(t *testing.T) {
	arr := []int{0, 0, 0, 3, 5, 5, 9, 9, 9, 9, 9, 9, 9, 9, 12}
	equal[int](t, "测试重复数组中目标,递归,存在,最左", SearchByRecursionLeftInRepeatArr(arr, 9), 6)
	equal[int](t, "测试重复数组中目标,递归,首个,存在,最左", SearchByRecursionLeftInRepeatArr(arr, 0), 0)
	equal[int](t, "测试重复数组中目标,递归,不存在,最左", SearchByRecursionLeftInRepeatArr(arr, 2), -1)

	equal[int](t, "测试重复数组中目标,循环,存在,最左", SearchByForLeftInRepeatArr(arr, 9), 6)
	equal[int](t, "测试重复数组中目标,循环,首个,存在,最左", SearchByForLeftInRepeatArr(arr, 0), 0)
	equal[int](t, "测试重复数组中目标,循环,不存在,最左", SearchByForLeftInRepeatArr(arr, 2), -1)

	equal[int](t, "测试重复数组中目标,循环,存在,最右", SearchByForRightInRepeatArr(arr, 9), 13)
	equal[int](t, "测试重复数组中目标,循环,首个,存在,最右", SearchByForRightInRepeatArr(arr, 0), 2)
	equal[int](t, "测试重复数组中目标,循环,不存在,最右", SearchByForRightInRepeatArr(arr, 2), -1)

}
