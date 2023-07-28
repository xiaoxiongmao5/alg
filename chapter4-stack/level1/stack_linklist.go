package level1
import (
	"errors"
)

func NewStack2[T any]() *Stack2[T] {
	return &Stack2[T]{}
}

type LinkNode[T any] struct {
	Val T 
	Next *LinkNode[T]
}
type Stack2[T any] struct {
	Head *LinkNode[T]
}

func (this *Stack2[T]) Push(val T) error {
	if this.IsEmpty() {
		this.Head = &LinkNode[T]{
			Val:val,
		}
		return nil
	}
	newNode := &LinkNode[T]{
		Val:val,
		Next:this.Head,
	}
	this.Head = newNode
	return nil
}

func (this *Stack2[T]) Pop() (val T, err error) {
	val, err = this.Peek()
	if err == nil {
		this.Head = this.Head.Next
	}
	return val, err
}

func (this *Stack2[T]) Peek() (val T, err error) {
	if this.IsEmpty() {
		return val, errors.New("stack empty")
	}
	return this.Head.Val, nil
}

func (this *Stack2[T]) IsEmpty() bool {
	return this.Head == nil
}

