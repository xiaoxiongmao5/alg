package stack

import (
	"fmt"
	"errors"
)

type Stack struct {
	MaxLen int
	Array [10]int
	Top int	//-1 为空
}

func (this *Stack) Pop() (int, error) {
	if this.IsEmpty() {
		return -1, errors.New("stack empty")
	}
	val := this.Array[this.Top]
	this.Top--
	return val, nil
}

func (this *Stack) Push(val int) error {
	if this.IsFull() {
		return errors.New("stack full")
	}
	this.Top++
	this.Array[this.Top] = val
	return nil
}

func (this *Stack) IsEmpty() bool {
	if this.Top == -1 {
		return true
	}
	return false
}

func (this *Stack) IsFull() bool {
	if this.Top == this.MaxLen - 1 {
		return true
	}
	return false
}

func (this *Stack) ShowStack() {
	if this.IsEmpty() {
		fmt.Println("stack empty")
		return
	}
	tmpNum := this.Top
	for {
		fmt.Printf("%d ==> ", this.Array[tmpNum])
		tmpNum--
		if tmpNum == -1 {
			break
		}
	}
	fmt.Println("")
}

func (this *Stack) GetStackSlice() []int {
	var ret = make([]int, 0)
	if this.IsEmpty() {
		fmt.Println("stack empty")
		return ret
	}
	// 注意：这里不是出栈顺序，而是从栈底开始的
	tmpNum := this.Top
	i := 0
	for {
		ret = append(ret, this.Array[i])
		i++
		if i > tmpNum {
			break
		}
	}
	return ret
}