package challenge

// 三指针法
func Method2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	slow, fast := 0, 0
	pre := nums[fast]-1
	for ; fast<length; pre,fast=nums[fast],fast+1 {
		if fast==length-1 && nums[fast]>pre {
			nums[slow] = nums[fast]
			slow++
		} else if nums[fast]>pre && nums[fast]<nums[fast+1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 双指针法
func Method(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	slow, fast := 0, 0
	for ; fast<length; fast++ {
		if fast<1 || nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
			continue
		} else {
			// 让fast走到不重复的位置
			for nums[fast] == nums[slow-1] {
				fast++
			}
			nums[slow-1] = nums[fast]
		}
	}
	return slow
}