package level1

// 题目链接：LeetCode-113. 路径总和 II：https://leetcode.cn/problems/path-sum-ii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	path := make([]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		// 叶子节点
		if node.Left == nil && node.Right == nil {
			// 路径匹配，加入结果列表
			if node.Val == sum {
				pathcopy := make([]int, len(path))
				copy(pathcopy, path)
				ret = append(ret, pathcopy)
			}
			path = path[:len(path)-1]
			return
		}
		dfs(node.Left, sum-node.Val)
		dfs(node.Right, sum-node.Val)
		path = path[:len(path)-1]
	}
	dfs(root, targetSum)
	return ret
}
