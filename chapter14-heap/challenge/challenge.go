package challenge

func getMaxK(arr []int, k int) int {
	length := len(arr)
	if length < k {
		return -1
	}
	// 生成一个长度为k的最小堆
	makeMinHeap(arr, k)
	for i := k; i < length; i++ {
		// 遇到比堆顶大的元素，替换掉堆顶，并最小化堆
		if arr[i] > arr[0] {
			arr[i], arr[0] = arr[0], arr[i]
			minHeap(arr, 0, k)
		}
	}
	return arr[0]
}

// 生成最小堆
func makeMinHeap(arr []int, length int) {
	// 最后一个非叶子结点
	for i := length/2 - 1; i >= 0; i-- {
		minHeap(arr, i, length)
	}
}

// 最小堆化
func minHeap(arr []int, i int, length int) {
	left, right, min := 2*i+1, 2*i+2, i
	if left < length && arr[left] < arr[min] {
		min = left
	}
	if right < length && arr[right] < arr[min] {
		min = right
	}
	if min != i {
		arr[min], arr[i] = arr[i], arr[min]
		minHeap(arr, min, length)
	}
}
