package level1

// 题目链接：LeetCode-77. 组合：https://leetcode.cn/problems/combinations/

func combine(n int, k int) [][]int {
	ret := make([][]int, 0)
	if k <= 0 || n < k {
		return ret
	}
	path := make([]int, 0)
	var dfs func(int)
	dfs = func(start int) {
		if len(path) == k {
			// 关键
			pathcopy := make([]int, k)
			copy(pathcopy, path)
			ret = append(ret, pathcopy)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return ret
}
