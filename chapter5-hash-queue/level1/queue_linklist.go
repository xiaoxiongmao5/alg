package level1

import (
	"errors"
	"fmt"
	"strings"
)

func NewQueue[T any]() *Queue[T]{
	return &Queue[T]{}
}

type Node[T any] struct {
	Val T
	Next *Node[T]
}

type Queue[T any] struct {
	Head *Node[T]
	Count int
}

func (this *Queue[T]) IsEmpty() bool {
	if this.Head == nil || this.Head.Next == nil {
		return true
	}
	return false
}

func (this *Queue[T]) Push(val T) {
	node := &Node[T]{Val:val}
	if this.IsEmpty() {
		this.Head = &Node[T]{
			Next:node,
		}
		this.Count++
		return
	}
	tmp := this.Head.Next
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = node
	this.Count++
	return
}

func (this *Queue[T]) Pop() (val T, err error) {
	if this.IsEmpty() {
		return val, errors.New("queue empty")
	}
	tmp := this.Head.Next
	this.Head.Next = tmp.Next
	this.Count--
	return tmp.Val, err
}

func (this *Queue[T]) Show() string {
	if this.IsEmpty() {
		return ""
	}
	tmp := this.Head.Next
	ret := make([]string, 0)
	for tmp != nil {
		ret = append(ret, fmt.Sprintf("%v", tmp.Val))
		tmp = tmp.Next
	}
	return strings.Join(ret, "=>")
}