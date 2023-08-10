package challenge

// 递归-二分查找
func Search(arr []int, val int) int {
	length := len(arr)

	find := -1
	var a func([]int, int, int, int)
	a = func(arr []int, left int, right int, val int) {
		if left > right {
			return
		}
		mid := (right-left)>>1 + left
		if arr[mid] == val {
			find = mid
			a(arr, 0, mid-1, val)
		} else if arr[mid] < val {
			a(arr, mid+1, right, val)
		} else {
			a(arr, 0, mid-1, val)
		}
	}
	a(arr, 0, length-1, val)
	return find
}

// 循环-二分查找
func Search2(arr []int, val int) int {
	length := len(arr)
	left, right := 0, length-1
	find := -1
	for left <= right {
		mid := (right-left)>>1 + left
		if arr[mid] == val {
			find = mid
			right = mid - 1
		} else if arr[mid] < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return find
}
