package challenge

func FindNoExistNumsBy1G(arr []int) (res []int) {
	N := 16 //分成16块
	for i := 0; i < N; i++ {
		res = append(res, FindNoExistNumsByPiece(arr, i)...)
	}
	return
}

// 单独统计每块中不存在的整数
func FindNoExistNumsByPiece(arr []int, targetIndex int) (res []int) {
	N := 268435456                //分成16块，每块处理的平均数
	bitmap := make([]int, N/32+1) //最多占用8MB

	// 统计当前块中数据的存在情况
	for _, num := range arr {
		if num%N == targetIndex {
			index := num / 32
			offset := num % 32
			mark := 1 << offset
			bitmap[index] |= mark
		}
	}
	// 找出不存在的数据
	for index, val := range bitmap {
		for i := 0; i < 32; i++ {
			mark := 1 << i
			if val&mark == 0 {
				// 计算不存在的整数值并添加到结果中
				num := i + index*32 + targetIndex*N
				res = append(res, num)
			}
		}
	}
	return
}
