package level1

// 二分查找-递归方式实现
func SearchByRecursion(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	return findTarget(nums, 0, length-1, target)
}

func findTarget(arr []int, left int, right int, target int) int {
	ret := -1
	if left > right {
		return ret
	}
	mid := left + ((right - left) >> 1)
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return findTarget(arr, mid+1, right, target)
	} else {
		return findTarget(arr, left, mid-1, target)
	}
}

// 二分查找-循环方式实现
func SearchByFor(nums []int, target int) int {
	ret := -1
	length := len(nums)
	if length == 0 {
		return ret
	}
	left, right := 0, length-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ret
}

// 二分查找-递归方式实现-数组中有重复的元素，找到最左边的
func SearchByRecursionLeftInRepeatArr(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	return findTargetRepeatArr(nums, 0, length-1, target)
}

func findTargetRepeatArr(arr []int, left int, right int, target int) int {
	ret := -1
	if left > right {
		return ret
	}
	mid := left + ((right - left) >> 1)
	if arr[mid] == target {
		for mid > 0 && arr[mid] == target {
			mid--
			if mid == 0 && arr[mid] == target {
				return mid
			}
		}
		return mid + 1
	} else if arr[mid] < target {
		return findTargetRepeatArr(arr, mid+1, right, target)
	} else {
		return findTargetRepeatArr(arr, left, mid-1, target)
	}
}

// 二分查找-循环方式实现-数组中有重复的元素，找到最左边的
func SearchByForLeftInRepeatArr(nums []int, target int) int {
	ret := -1
	length := len(nums)
	if length == 0 {
		return ret
	}
	left, right := 0, length-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			for mid > 0 && nums[mid] == target {
				mid--
				if mid == 0 && nums[mid] == target {
					return mid
				}
			}
			return mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ret
}

// 二分查找-循环方式实现-数组中有重复的元素，找到最右边的
func SearchByForRightInRepeatArr(nums []int, target int) int {
	ret := -1
	length := len(nums)
	if length == 0 {
		return ret
	}
	left, right := 0, length-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			for mid < length && nums[mid] == target {
				mid++
				if mid == length-1 && nums[mid] == target {
					return mid
				}
			}
			return mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ret
}
