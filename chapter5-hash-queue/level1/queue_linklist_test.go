package level1

import (
	"testing"
	"fmt"
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

func TestNewQueue(t *testing.T) {
	queue := NewQueue[int]()
	queue.Pop()
	equal(t, "空队列Pop", queue.Show(), "")

	queue.Push(1)
	equal(t, "空队列Push", queue.Show(), "1")

	queue.Pop()
	equal(t, "队列Pop", queue.Show(), "")

	queue.Push(2)
	equal(t, "队列Push", queue.Show(), "2")

	queue.Push(3)
	equal(t, "队列Push", queue.Show(), "2=>3")

	queue.Push(4)
	equal(t, "队列Push", queue.Show(), "2=>3=>4")

	queue.Pop()
	equal(t, "队列Pop", queue.Show(), "3=>4")

	queue.Pop()
	equal(t, "队列Pop", queue.Show(), "4")

	queue.Pop()
	equal(t, "队列Pop", queue.Show(), "")
}