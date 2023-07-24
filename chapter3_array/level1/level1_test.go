package level1

import (
	"testing"
	"fmt"
)

func ShowArr2Str(arr []int) string {
	return fmt.Sprintf("%v", arr)
}

func equal(t *testing.T, desc, arrstr1, arrstr2 string) {
	fmt.Printf("======= %s =======\n", desc)
	if arrstr1 == arrstr2 {
		// fmt.Printf("成功\n export\t%s \n got\t%s \n", arrstr2, arrstr1)
	} else {
		t.Errorf("失败\n export\t%s \n got\t%s \n", arrstr2, arrstr1)
	}
	fmt.Println("")
}

func TestAddFirstInArr(t *testing.T) {
	fmt.Printf("======= 【【 测试：%s 】】 =======\n", "添加数组-首位元素")
	equal(t, "在数组首部添加元素", ShowArr2Str(AddFirstInArr([]int{1,2,3,4,5}, 7)), ShowArr2Str([]int{7,1,2,3,4,5}))

	equal(t, "在空数组首部添加元素", ShowArr2Str(AddFirstInArr([]int{}, 7)), ShowArr2Str([]int{7}))
}

func TestAddLastInArr(t *testing.T) {
	fmt.Printf("======= 【【 测试：%s 】】 =======\n", "添加数组-末尾元素")
	equal(t, "在数组尾部添加元素", ShowArr2Str(AddLastInArr([]int{1,2,3,4,5}, 7)), ShowArr2Str([]int{1,2,3,4,5,7}))

	equal(t, "在空数组尾部添加元素", ShowArr2Str(AddLastInArr([]int{}, 7)), ShowArr2Str([]int{7}))
}

func TestAddInArr(t *testing.T) {
	fmt.Printf("======= 【【 测试：%s 】】 =======\n", "添加数组-中间元素")
	equal(t, "在数组中间添加元素", ShowArr2Str(AddInArr([]int{1,2,3,4,5}, 6)), ShowArr2Str([]int{1,2,3,4,5,6}))

	equal(t, "在数组中间添加元素", ShowArr2Str(AddInArr([]int{9,7,6,4,2,1}, 4)), ShowArr2Str([]int{9,7,6,4,4,2,1}))
}

func TestDelFirstInArr(t *testing.T) {
	fmt.Printf("======= 【【 测试：%s 】】 =======\n", "删除数组-首位元素")
	equal(t, "删除首位", ShowArr2Str(DelFirstInArr([]int{1,2,3,4,5})), ShowArr2Str([]int{2,3,4,5}))

	equal(t, "删除空数组首位", ShowArr2Str(DelFirstInArr([]int{})), ShowArr2Str([]int{}))
}

func TestDelLastInArr(t *testing.T) {
	fmt.Printf("======= 【【 测试：%s 】】 =======\n", "删除数组-末尾元素")
	equal(t, "删除最后一个", ShowArr2Str(DelLastInArr([]int{1,2,3,4,5})), ShowArr2Str([]int{1,2,3,4}))

	equal(t, "删除空数组最后一个", ShowArr2Str(DelLastInArr([]int{})), ShowArr2Str([]int{}))
}

func TestDelInArr(t *testing.T) {
	fmt.Printf("======= 【【 测试：%s 】】 =======\n", "删除数组-中间元素")
	equal(t, "中间删除存在的", ShowArr2Str(DelInArr([]int{1,2,3,4,5}, 2)), ShowArr2Str([]int{1,3,4,5}))

	equal(t, "中间删除不存在的", ShowArr2Str(DelInArr([]int{1,2,3,4,5}, 6)), ShowArr2Str([]int{1,2,3,4,5}))

	equal(t, "空数组删除", ShowArr2Str(DelInArr([]int{}, 4)), ShowArr2Str([]int{}))
}


// func TestA(t *testing.T) {
// 	A()
// }