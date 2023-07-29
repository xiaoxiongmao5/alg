package level2

import (
	"errors"
)

func NewMaxStack[T int|int64](maxNum int) *MaxStack[T] {
	return &MaxStack[T] {
		MaxNum:maxNum,
		Array:make([]T, maxNum),
		MaxArray:make([]T, maxNum),
		TopNum:0,
	}
}

type MaxStack[T int|int64] struct {
	MaxNum int
	Array []T 
	MaxArray []T 
	TopNum int	//栈顶从0开始
}

// 将元素压入栈中
func (this *MaxStack[T]) Push(val T) error {
	if this.IsFull() {
		this.ExpandCap(this.MaxNum)
	}
	if this.IsEmpty() {
		this.MaxArray[this.TopNum] = val
	} else {
		preMaxVal := this.MaxArray[this.TopNum-1]
		if val >= preMaxVal {
			this.MaxArray[this.TopNum] = val
		} else {
			this.MaxArray[this.TopNum] = preMaxVal
		}
	}
	this.Array[this.TopNum] = val
	this.TopNum++
	return nil
}

// 移除栈顶元素并返回这个元素
func (this *MaxStack[T]) Pop() (val T, err error) {
	val, err = this.Top()
	if err != nil {
		return val, err
	}
	this.TopNum--
	return val, err
}

func (this *MaxStack[T]) IsEmpty() bool {
	return this.TopNum == 0
}

func (this *MaxStack[T]) IsFull() bool {
	return this.TopNum == this.MaxNum
}

func (this *MaxStack[T]) ExpandCap(size int) {
	tmpArray := make([]T, size)
	this.Array = append(this.Array, tmpArray ...)
	this.MaxArray = append(this.MaxArray, tmpArray ...)
	this.MaxNum = this.MaxNum+size
}

// 返回栈顶元素，无需移除
func (this *MaxStack[T]) Top() (val T, err error){
	if this.IsEmpty() {
		return val, errors.New("stack empty")
	}
	val = this.Array[this.TopNum-1]
	return
}

// 检索并返回栈中最大元素，无需移除
func (this *MaxStack[T]) PeekMax() (val T, err error) {
	if this.IsEmpty() {
		return val, errors.New("stack empty")
	}
	val = this.MaxArray[this.TopNum-1]
	return
}

// 检索并返回栈中最大元素，并将其移除
func (this *MaxStack[T]) PopMax() (val T, err error) {
	val, err = this.PeekMax()
	if err != nil {
		return val, err
	}
	tmpArray := make([]T, 0)
	
	for {
		tmpV, tmpErr := this.Pop()
		if tmpErr != nil {
			break
		}
		if tmpV == val {
			break
		}
		tmpArray = append(tmpArray, tmpV)
	}
	for len(tmpArray) != 0 {
		tmpLen := len(tmpArray)
		this.Push(tmpArray[tmpLen-1])
		tmpArray = tmpArray[:tmpLen-1]
	}
	return val ,err
}

func (this *MaxStack[T]) GetValidArray() []T {
	retArr := make([]T, this.MaxNum)
	for i:=0; i<this.TopNum; i++ {
		retArr[i] = this.Array[i]
	}
	return retArr
}