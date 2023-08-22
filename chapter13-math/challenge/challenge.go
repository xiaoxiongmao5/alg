package challenge

import (
	"fmt"
	"math"
)

func Method(num int) (res int) {
	flag := false //是否进位-1标识
	for i := 0; num > 0; i++ {
		n := num % 10 //余数
		if n == 4 {
			return -1
		}
		if !flag {
			n--
			if n == 4 {
				n-- //跳过4
			}
			if n < 0 {
				n = 9
				flag = false
			} else {
				flag = true
			}
		}
		res = n*int(math.Pow(10, float64(i))) + res
		num = num / 10
	}
	return
}

func equal2[T string | int](str1 T, str2 T) {
	if str1 == str2 {
		fmt.Printf("\n通过\n export\t%v\n got\t%v\n", str2, str1)
	} else {
		fmt.Printf("\n未通过\n export\t%v\n got\t%v\n", str2, str1)
	}
	fmt.Println("")
}

func main() {
	equal2[int](Method(5), 3)
	equal2[int](Method(15), 13)
	equal2[int](Method(25), 23)
	equal2[int](Method(50), 39)
	equal2[int](Method(500), 399)
	equal2[int](Method(5000), 3999)
}
