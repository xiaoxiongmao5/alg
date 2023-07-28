package level1

import (
	"errors"
)

func NewStack1[T any](maxNum int) *Stack1[T]{
	slice := make([]T, maxNum)
	return &Stack1[T]{
		MaxNum:maxNum,
		Array:slice,
		Top:0,
	}
}

type Stack1[T any] struct {
	MaxNum int
	Array []T
	Top int	//栈顶上的位置
}

func (this *Stack1[T]) Push(val T) error {
	if this.IsFull() {
		this.ExpandCap(this.MaxNum)
	}
	this.Array[this.Top] = val
	this.Top++
	return nil
}

func (this *Stack1[T]) Pop() (val T, err error) {
	if this.IsEmpty() {
		return val, errors.New("stack empty")
	}
	this.Top--
	return this.Array[this.Top], nil
}

func (this *Stack1[T]) Peek() (val T, err error) {
	if this.IsEmpty() {
		return val, errors.New("stack empty")
	}
	return this.Array[this.Top-1], nil
}

func (this *Stack1[T]) IsFull() bool {
	return this.MaxNum == this.Top
}

func (this *Stack1[T]) IsEmpty() bool {
	return this.Top == 0
}

func (this *Stack1[T]) ExpandCap(size int) {
	tmpArray := make([]T, this.MaxNum+size)
	this.Array = append(this.Array, tmpArray ...)
	this.MaxNum = this.MaxNum+size
}