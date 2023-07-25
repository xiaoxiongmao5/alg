package challenge

func Method(nums []int) int {
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