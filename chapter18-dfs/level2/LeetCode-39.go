package level2

import (
	"strconv"
	"strings"
)

func combinationSum(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	maps := make(map[string]bool, 0)
	var dfs func(int)
	dfs = func(tar int) {
		for _, v := range candidates {
			if v > tar {
				continue
			}
			path = append(path, v)
			if v == tar {
				pathcopy := make([]int, len(path))
				copy(pathcopy, path)
				sort(pathcopy)
				key := getKeyStr(pathcopy)
				if _, ok := maps[key]; !ok {
					maps[key] = true
					ret = append(ret, pathcopy)
				}
				path = path[:len(path)-1]
				continue
			}
			dfs(tar - v)
			path = path[:len(path)-1]
		}
	}
	dfs(target)
	return ret
}
func getKeyStr(arr []int) string {
	builder := strings.Builder{}
	for _, v := range arr {
		builder.WriteString(strconv.Itoa(v))
	}
	return builder.String()
}
func sort(arr []int) {
	length := len(arr)
	if length <= 0 {
		return
	}
	sortQuick(arr, 0, length-1)
}
func sortQuick(arr []int, start, end int) {
	if start >= end {
		return
	}
	midVal := arr[(end-start)>>1+start]
	left, right := start, end
	for left <= right {
		for left <= right && arr[left] < midVal {
			left++
		}
		for left <= right && arr[right] > midVal {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	sortQuick(arr, start, right)
	sortQuick(arr, left, end)
}
