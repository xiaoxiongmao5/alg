package challenge

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    ret_arr := make([]int, 0)
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        length := len(queue)
        ret_arr = append(ret_arr, queue[0].Val)
        for i:=0;i<length;i++ {
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
        }
        queue = queue[length:]
    }
    return ret_arr
}