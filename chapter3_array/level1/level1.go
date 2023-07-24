package level1

import (
	"fmt"
)

// 在数组首部插入元素
func AddFirstInArr(arr []int, val int) []int{
	newArr := make([]int, len(arr)+1)
	newArr[0] = val
	// 下面for循环也可以换成内置copy函数： 
	// copy(newArr[1:], arr)
	for i:=1;i<len(newArr);i++ {
		newArr[i] = arr[i-1]
	}
	return newArr
}

// 在数组尾部插入元素
func AddLastInArr(arr []int, val int) []int{
	newArr := make([]int, len(arr)+1)
	// 下面for循环也可以换成内置copy函数： 
	// copy(newArr, arr)
	for i:=0;i<len(arr);i++ {
		newArr[i] = arr[i]
	}
	newArr[len(newArr)-1] = val
	return newArr
}

// 在数组中间插入元素
func AddInArr(arr []int, val int) []int{
	arrlen := len(arr)
	left,right := 0, arrlen-1
	// 假设应该插到最后
	findindex := arrlen
	// 使用二分查找法找到val对应的位置
	// 升序
	if arr[left] <= arr[right] {
		for left <= right {
			mid := (right-left)/2+left
			if val <= arr[mid] {
				right = mid-1
				findindex = mid
			} else {
				left = mid+1
			}
		}
	} else {
		for left <= right {
			mid := (right-left)/2+left
			if val >= arr[mid] {
				right = mid-1
				findindex = mid
			} else {
				left = mid+1
			}
		}
	}
	newArr := make([]int, arrlen+1)
	// 下面for循环也可以换成内置copy函数： 
	// copy(newArr[:findindex], arr)
	// copy(newArr[findindex+1:], arr[findindex:])
	for i:=0;i<arrlen;i++ {
		if i < findindex {
			newArr[i] = arr[i]
		} else {
			newArr[i+1] = arr[i]
		}
	}
	newArr[findindex] = val
	return newArr
}

// 删除数组首位
func DelFirstInArr(arr []int) []int{
	if len(arr) == 0 {
		return arr
	}
	return arr[1:]
}

// 删除数组末尾元素
func DelLastInArr(arr []int) []int{
	if len(arr) == 0 {
		return arr
	}
	return arr[:len(arr)-1]
}
// 删除数组中间元素
func DelInArr(arr []int, val int) []int{
	arrlen := len(arr)
	if arrlen == 0 {
		return arr
	}
	findindex := -1
	for i, k := range arr {
		if k == val {
			findindex = i
			break
		}
	}
	if findindex == -1 {
		fmt.Println("数组中不存在要删除的元素 ", val)
		return arr
	}
	// 将要删除元素后的所有元素都往前移动一位
	copy(arr[findindex:], arr[findindex+1:])
	return arr[:arrlen-1]
}


func A() {
	fmt.Println("slice := arr[:3]")
	arr := [6]int{1,2,3,4,5,6}
	slice := arr[:3]
	fmt.Println("arr=",arr)
	fmt.Println("slice=",slice)
	fmt.Println("")

	fmt.Println("slice[0] = 10")
	slice[0] = 10
	fmt.Println("arr=",arr)
	fmt.Println("slice=",slice)
	fmt.Println("")

	fmt.Println("arr[1] = 11")
	arr[1] = 11
	fmt.Println("arr=",arr)
	fmt.Println("slice=",slice)
	fmt.Println("")

	fmt.Println("arr[5] = 12")
	arr[5] = 12
	fmt.Println("arr=",arr)
	fmt.Println("slice=",slice)
	fmt.Println("")

	fmt.Println("slice2 := slice[1:] slice2[0] = 40")
	slice2 := slice[1:]
	slice2[0] = 40
	fmt.Println("arr=",arr)
	fmt.Println("slice=",slice)
	fmt.Println("slice2=",slice2)
	fmt.Println("")

	fmt.Println("slice = append(slice, 50)")
	slice = append(slice, 50)
	fmt.Println("arr=",arr)
	fmt.Println("slice=",slice)
	fmt.Println("slice2=",slice2)
	fmt.Println("")

	fmt.Println("slice2 = append(slice2, 60)")
	slice2 = append(slice2, 60)
	fmt.Println("arr=",arr)
	fmt.Println("slice=",slice)
	fmt.Println("slice2=",slice2)
	fmt.Println("")

	/**
	从以上测试可以发现：
	通过切片方式从数组上获得的切片，那这个切片的底层数组就是这个数组，也就是说之后对该切片做了任何赋值、append的操作都会影响到原来的数组。
	即使从这个切片上再切片出来的切片，修改也是一样的同步到底层数组以及其所有的引用切片上。
	append是为了扩容，底层原理是新构造一个数组，然后将原来数组的值都拷贝过来，然后将所有的引用也同步到新的数组，之后的修改，不仅会影响所有切片，原始的数组也一样受影响
	copy只是做值拷贝，不会影响原数组或切片
	*/


	fmt.Println("slice3 := make([]int, 6)  copy(slice3, slice)")
	slice3 := make([]int, 6)
	copy(slice3, slice)
	fmt.Println("arr=",arr)
	fmt.Println("slice=",slice)
	fmt.Println("slice2=",slice2)
	fmt.Println("slice3=",slice3)
	fmt.Println("")

	fmt.Println("slice3[0] = 80")
	slice3[0] = 80
	fmt.Println("arr=",arr)
	fmt.Println("slice=",slice)
	fmt.Println("slice2=",slice2)
	fmt.Println("slice3=",slice3)
	fmt.Println("")
	// arr := make([]int, 7 ,10)
	// arr[0] = 0
	// arr[1] = 1
	// arr[2] = 2
	// arr[3] = 3
	// fmt.Printf("arr=%#v \n", arr)
	// fmt.Printf("arr.len=%d \n", len(arr))
	// fmt.Printf("arr.cap=%d \n", cap(arr))
	// fmt.Println("")
	// arr[4] = 4
	// fmt.Printf("arr=%#v \n", arr)
	// fmt.Printf("arr.len=%d \n", len(arr))
	// fmt.Printf("arr.cap=%d \n", cap(arr))
	// fmt.Println("")
	// arr[6] = 6
	// fmt.Printf("arr=%#v \n", arr)
	// fmt.Printf("arr.len=%d \n", len(arr))
	// fmt.Printf("arr.cap=%d \n", cap(arr))
	// fmt.Println("")

	// arr = arr[:len(arr)-1]
	// fmt.Printf("arr=%#v \n", arr)
	// fmt.Printf("arr.len=%d \n", len(arr))
	// fmt.Printf("arr.cap=%d \n", cap(arr))
	// fmt.Println("")

	// newArr := make([]int, 10, 11)
	// copy(newArr, arr[:len(arr)-1])
	// fmt.Printf("newArr=%#v \n", newArr)
	// fmt.Printf("newArr.len=%d \n", len(newArr))
	// fmt.Printf("newArr.cap=%d \n", cap(newArr))
	// fmt.Println("")
}