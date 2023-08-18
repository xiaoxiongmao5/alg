package challenge

import (
	"fmt"
	"strings"
)

func Compress(arr []string) {
	for i, str := range arr {
		left, right := 0, 0
		length := len(str)
		// 如果是奇数，不可能重复
		if length%2 != 0 {
			continue
		}
		// 让left走到中间
		for right < length {
			left++
			right = right + 2
		}
		// 判断是否是重复的
		isrepeat := true
		start := 0
		for left < length {
			if str[start] != str[left] {
				isrepeat = false
				break
			}
			start++
			left++
		}
		if !isrepeat {
			continue
		}
		arr[i] = str[:start] + "~"
	}
}

func main() {
	arr := []string{"tiantian", "yuanyuan", "tingting", "yingzi", "qiyue"}
	compareArr := []string{"tian~", "yuan~", "ting~", "yingzi", "qiyue"}
	Compress(arr)
	str1 := strings.Join(arr, ",")
	str2 := strings.Join(compareArr, ",")
	if str1 == str2 {
		fmt.Printf("\n通过\n export\t%v\n got\t%v\n", str2, str1)
	} else {
		fmt.Printf("\n未通过\n export\t%v\n got\t%v\n", str2, str1)
	}
}
