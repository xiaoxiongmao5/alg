package challenge

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DeleteNode(root *TreeNode) *TreeNode {
	if deleteLeaf(root) {
		return nil
	}
	return root
}

func deleteLeaf(node *TreeNode) bool {
	if node == nil {
		return true
	}
	if node.Left == nil && node.Right == nil {
		return true
	}
	if deleteLeaf(node.Left) {
		node.Left = nil
	}
	if deleteLeaf(node.Right) {
		node.Right = nil
	}
	return false
}

func DeleteNode1(root *TreeNode) (ret *TreeNode) {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return nil
	}
	a(root.Left, root, "left")
	a(root.Right, root, "right")
	return root
}

func a(root *TreeNode, father *TreeNode, flag string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if flag == "left" {
			father.Left = nil
		} else if flag == "right" {
			father.Right = nil
		}
		return
	}
	a(root.Left, root, "left")
	a(root.Right, root, "right")
}
