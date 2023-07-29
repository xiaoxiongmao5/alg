package level2

import (
	"testing"
	"fmt"
	"strings"
)

func equal(t *testing.T, desc string, str1 string, str2 string) {
	fmt.Printf("======= %s =======\n", desc)
	if str1 == str2 {
		fmt.Printf("\n通过\n export\t%v\n got\t%v\n", str2, str1)
	} else {
		t.Errorf("\n未通过\n export\t%v\n got\t%v\n", str2, str1)
	}
	fmt.Println("")
}

func arr2str(arr []int) string{
	length := len(arr)
	ret := make([]string, length)
	for i:=0;i<length; i++ {
		ret[i] = fmt.Sprintf("%d", arr[i])
	}
	return strings.Join(ret, "=>")
}

func TestNewMaxStack(t *testing.T) {
	stack := NewMaxStack[int](5)
	stack.Push(5)
	equal(t, "Push元素", arr2str(stack.GetValidArray()), arr2str([]int{5, 0, 0, 0, 0}))
	
	stack.Push(1)
	equal(t, "Push元素", arr2str(stack.GetValidArray()), arr2str([]int{5, 1, 0, 0, 0}))

	stack.Push(5)
	equal(t, "Push元素", arr2str(stack.GetValidArray()), arr2str([]int{5, 1, 5, 0, 0}))

	val, _ := stack.Top()
	equal(t, "Top:返回栈顶元素，无需移除", arr2str(stack.GetValidArray()), arr2str([]int{5, 1, 5, 0, 0}))
	equal(t, "Top:返回栈顶元素，无需移除", fmt.Sprintf("%d", val), "5")

	val, _ = stack.PopMax()
	equal(t, "PopMax:检索并返回栈中最大元素，并将其移除", arr2str(stack.GetValidArray()), arr2str([]int{5, 1, 0, 0, 0}))
	equal(t, "PopMax:检索并返回栈中最大元素，并将其移除", fmt.Sprintf("%d", val), "5")

	val, _ = stack.Top()
	equal(t, "Top:返回栈顶元素，无需移除", arr2str(stack.GetValidArray()), arr2str([]int{5, 1, 0, 0, 0}))
	equal(t, "Top:返回栈顶元素，无需移除", fmt.Sprintf("%d", val), "1")

	val, _ = stack.PeekMax()
	equal(t, "PeekMax:检索并返回栈中最大元素，无需移除", arr2str(stack.GetValidArray()), arr2str([]int{5, 1, 0, 0, 0}))
	equal(t, "PeekMax:检索并返回栈中最大元素，无需移除", fmt.Sprintf("%d", val), "5")

	val, _ = stack.Pop()
	equal(t, "Pop:移除栈顶元素并返回这个元素", arr2str(stack.GetValidArray()), arr2str([]int{5, 0, 0, 0, 0}))
	equal(t, "Pop:移除栈顶元素并返回这个元素", fmt.Sprintf("%d", val), "1")

	val, _ = stack.Top()
	equal(t, "Top:返回栈顶元素，无需移除", arr2str(stack.GetValidArray()), arr2str([]int{5, 0, 0, 0, 0}))
	equal(t, "Top:返回栈顶元素，无需移除", fmt.Sprintf("%d", val), "5")
}