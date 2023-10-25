package level1

import (
	"strconv"
	"strings"
)

// 题目链接：LeetCode-257. 二叉树的所有路径：https://leetcode.cn/problems/binary-tree-paths/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	ret := make([]string, 0)
	if root == nil {
		return ret
	}
	path := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			ret = append(ret, conv(path))
			path = path[:len(path)-1]
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	return ret
}
func conv(arr []int) string {
	length := len(arr)
	strarr := make([]string, length)
	for i, v := range arr {
		strarr[i] = strconv.Itoa(v)
	}
	return strings.Join(strarr, "->")
}
