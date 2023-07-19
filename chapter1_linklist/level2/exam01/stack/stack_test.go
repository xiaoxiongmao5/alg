package stack

import(
	"fmt"
	"testing"
	"strings"
	"reflect"
)

func helperCompare(slice []int) string {
	var ret = make([]string, 0)
	for _, v := range slice {
		ret = append(ret, fmt.Sprintf("%d", v))
	}
	return strings.Join(ret, "=>")
}

func equal(t *testing.T, desc string, slice1 []int, slice2 []int) {
	result := reflect.DeepEqual(slice1, slice2)
	if result {
		fmt.Printf("%s, 正常：%s \n", desc, helperCompare(slice1))
	} else {
		t.Errorf("%s, 错误\n except\t%s \n got\t%s \n", desc, helperCompare(slice2), helperCompare(slice1))
	}
}

func TestPush(t *testing.T) {
	fmt.Printf("======= %s =======\n", "测试可以正确插入数据")
	stack := &Stack{
		MaxLen:10,
		Top:-1,
	}
	for i:=0;i<10;i++{
		err := stack.Push(i)
		if err != nil {
			fmt.Println("err: ", err)
		}
	}
	stack.ShowStack()
	compare_slice := []int{0,1,2,3,4,5,6,7,8,9}
	equal(t, "栈内数据", stack.GetStackSlice(), compare_slice)

	fmt.Printf("======= %s =======\n", "测试栈满，不能插入")
	err := stack.Push(11)
	if err != nil {
		fmt.Println("err: ", err)
	}
	compare_slice = []int{0,1,2,3,4,5,6,7,8,9}
	equal(t, "栈内数据", stack.GetStackSlice(), compare_slice)
	
	fmt.Printf("======= %s =======\n", "测试栈的存取")
	val, _ := stack.Pop()
	equal(t, "pop出的数据", []int{val.(int)}, []int{9})
	val, _ = stack.Pop()
	equal(t, "pop出的数据", []int{val.(int)}, []int{8})
	val, _ = stack.Pop()
	equal(t, "pop出的数据", []int{val.(int)}, []int{7})
	val, _ = stack.Pop()
	equal(t, "pop出的数据", []int{val.(int)}, []int{6})
	val, _ = stack.Pop()
	equal(t, "pop出的数据", []int{val.(int)}, []int{5})
	val, _ = stack.Pop()
	equal(t, "pop出的数据", []int{val.(int)}, []int{4})
	val, _ = stack.Pop()
	equal(t, "pop出的数据", []int{val.(int)}, []int{3})
	val, _ = stack.Pop()
	equal(t, "pop出的数据", []int{val.(int)}, []int{2})
	compare_slice = []int{0,1}
	equal(t, "栈内数据", stack.GetStackSlice(), compare_slice)

	fmt.Printf("======= %s =======\n", "测试栈空，不能取出")
	val, _ = stack.Pop()
	equal(t, "pop出的数据", []int{val.(int)}, []int{1})
	val, _ = stack.Pop()
	equal(t, "pop出的数据", []int{val.(int)}, []int{0})
	val, err = stack.Pop()
	if err != nil {
		fmt.Println("err: ", err)
	}
	compare_slice = []int{}
	equal(t, "栈内数据", stack.GetStackSlice(), compare_slice)

	fmt.Printf("======= %s =======\n", "测试栈的存取")
	stack.Push(100)
	stack.Push(200)
	compare_slice = []int{100, 200}
	equal(t, "栈内数据,加100,再加200", stack.GetStackSlice(), compare_slice)
	stack.Pop()
	compare_slice = []int{100}
	equal(t, "栈内数据,减一位", stack.GetStackSlice(), compare_slice)
}