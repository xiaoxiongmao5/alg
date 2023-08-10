package challenge

func Search(arr []int, val int) int {
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
