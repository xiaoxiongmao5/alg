package level1

import (
	"testing"
	"fmt"
)

func TestNewStack2(t *testing.T) {
	stack2 := NewStack2[int]()
	_, err := stack2.Pop()
	fmt.Printf("======= %s =======\n","空栈取值")
	if err != nil {
		fmt.Println("成功。err=", err)
	} else {
		t.Errorf("失败")
	}

	fmt.Printf("======= %s =======\n","空栈存值")
	err = stack2.Push(10)
	if err == nil {
		fmt.Println("成功。")
	} else {
		t.Errorf("失败")
	}

	err = stack2.Push(20)
	if err == nil {
		fmt.Println("成功。")
	} else {
		t.Errorf("失败")
	}

	fmt.Printf("======= %s =======\n","栈取值")
	val, err := stack2.Pop()
	if err == nil && val == 20 {
		fmt.Printf("成功。\nexport:\t%d \ngot:\t%d\n", 20, val)
	} else {
		t.Errorf("失败。\nexport:\t%d \ngot:\t%d\n", 20, val)
	}

	val, err = stack2.Pop()
	if err == nil && val == 10 {
		fmt.Printf("成功。\nexport:\t%d \ngot:\t%d\n", 10, val)
	} else {
		t.Errorf("失败。\nexport:\t%d \ngot:\t%d\n", 10, val)
	}

	val, err = stack2.Pop()
	if err == nil && val == 10 {
		t.Errorf("失败。\nexport:\t%v \ngot:\t%v\n", "stack empty", err)
	} else {
		fmt.Printf("成功。\nexport:\t%v \ngot:\t%v\n", "stack empty", err)
	}

	fmt.Printf("======= %s =======\n","空栈存值")
	err = stack2.Push(10)
	if err == nil {
		fmt.Println("成功。")
	} else {
		t.Errorf("失败")
	}

	err = stack2.Push(20)
	if err == nil {
		fmt.Println("成功。")
	} else {
		t.Errorf("失败")
	}

	err = stack2.Push(30)
	if err == nil {
		fmt.Println("成功。")
	} else {
		t.Errorf("失败")
	}

	err = stack2.Push(40)
	if err == nil {
		fmt.Println("成功。")
	} else {
		t.Errorf("失败")
	}

	err = stack2.Push(50)
	if err == nil {
		fmt.Println("成功。")
	} else {
		t.Errorf("失败")
	}

	err = stack2.Push(60)
	if err == nil {
		fmt.Println("成功。")
	} else {
		t.Errorf("失败")
	}

	err = stack2.Push(70)
	if err == nil {
		fmt.Println("成功。")
	} else {
		t.Errorf("失败")
	}

	fmt.Printf("======= %s =======\n","栈看peek值")
	val, err = stack2.Peek()
	if err == nil && val == 70 {
		fmt.Printf("成功。\nexport:\t%d \ngot:\t%d\n", 70, val)
	} else {
		t.Errorf("失败。\nexport:\t%d \ngot:\t%d\n", 70, val)
	}

	fmt.Printf("======= %s =======\n","栈取值")
	val, err = stack2.Pop()
	if err == nil && val == 70 {
		fmt.Printf("成功。\nexport:\t%d \ngot:\t%d\n", 70, val)
	} else {
		t.Errorf("失败。\nexport:\t%d \ngot:\t%d\n", 70, val)
	}

	fmt.Printf("======= %s =======\n","栈看peek值")
	val, err = stack2.Peek()
	if err == nil && val == 60 {
		fmt.Printf("成功。\nexport:\t%d \ngot:\t%d\n", 60, val)
	} else {
		t.Errorf("失败。\nexport:\t%d \ngot:\t%d\n", 60, val)
	}

	fmt.Printf("======= %s =======\n","栈取值")
	val, err = stack2.Pop()
	if err == nil && val == 60 {
		fmt.Printf("成功。\nexport:\t%d \ngot:\t%d\n", 60, val)
	} else {
		t.Errorf("失败。\nexport:\t%d \ngot:\t%d\n", 60, val)
	}

	fmt.Printf("======= %s =======\n","栈看peek值")
	val, err = stack2.Peek()
	if err == nil && val == 50 {
		fmt.Printf("成功。\nexport:\t%d \ngot:\t%d\n", 50, val)
	} else {
		t.Errorf("失败。\nexport:\t%d \ngot:\t%d\n", 50, val)
	}
}