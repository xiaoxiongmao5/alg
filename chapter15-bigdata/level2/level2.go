package level2

func FindNoExistNumBy1G(arr []int) int {
	N := 4000000000
	bitmap := make([]int, N/32+1)
	for _, num := range arr {
		num0 := num - 1
		index := num0 / 32
		offset := num0 % 32
		mark := 1 << offset
		bitmap[index] |= mark
	}
	for index, v := range bitmap {
		for i := 0; i < 32; i++ {
			mark := 1 << i
			if v&mark == 0 {
				return index*32 + i + 1
			}
		}
	}
	return -1
}

func FindTargetIndex(arr []int) (targerIndex int) {
	N := 67108864
	countArr := make([]int, 64) //大约占用 0.25K 内存空间
	// 遍历40亿数据，明确每个数据对应的块，统计落在每个块上的总数量
	for _, num := range arr {
		countArr[num/N]++
	}
	for i, count := range countArr {
		if count < N {
			return i
		}
	}
	return -1
}
func FindNoExistNumBy10M(arr []int) (res int) {
	N := 67108864
	targetIndex := FindTargetIndex(arr)
	bitmap := make([]int, N/32+1) //大约占用 8M 内存空间
	for _, num := range arr {
		// 如果数据是属于目标块的
		if num/N == targetIndex {
			val := num - targetIndex*N
			val0 := val - 1
			index := val0 / 32
			offset := val0 % 32
			mark := 1 << offset
			bitmap[index] |= mark
		}
	}
	for index, intval := range bitmap {
		for i := 0; i < 32; i++ {
			mark := 1 << i
			// 找到不存在的值
			if intval&mark == 0 {
				res = i + index*32 + targetIndex*N + 1
				return res
			}
		}
	}
	return -1
}
