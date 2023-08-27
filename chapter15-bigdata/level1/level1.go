package level1

func FindDuplicatesByByteIn32000(arr []int) (duplicate []int) {
	// 32000个元素
	// 4KB = 4*1024(byte) = 4*1024*8(bit) = 32768(bit) > 32000
	// 所以可以使用 32000(bit) 存储 这32000个元素的存在状态
	N := 32000
	bitmap := make([]byte, N/8+1)
	for _, num := range arr {
		// 计算 num 在 bitmap 中的索引
		// index := num / 8
		index := num >> 3
		// 计算 num 在 bitmap 中的偏移量
		offset := num % 8
		mark := byte(1 << offset)
		if bitmap[index]&mark != 0 {
			duplicate = append(duplicate, num)
		} else {
			bitmap[index] |= mark
		}
	}
	return
}

func FindDuplicatesByIntIn32000(arr []int) (duplicate []int) {
	N := 32000
	// 或者这里不用+1，只要索引是base0的即可
	bitmap := make([]int, N/32)
	for _, num := range arr {
		num0 := num - 1 //base0开始
		index := num0 / 32
		offset := num0 % 32
		mark := 1 << offset
		if bitmap[index]&mark != 0 {
			duplicate = append(duplicate, num)
		} else {
			bitmap[index] |= mark
		}
	}
	return
}
