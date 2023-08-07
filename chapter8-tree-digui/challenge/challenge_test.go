package challenge

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func equal(t *testing.T, desc string, str1slice []string, str2slice []string) {
	fmt.Printf("======= %s =======\n", desc)
	var str1, str2 string
	str1 = strings.Join(str1slice, ",")
	str2 = strings.Join(str2slice, ",")
	if str1 == str2 {
		fmt.Printf("\n通过\n export\t%v\n got\t%v\n", str2, str1)
	} else {
		t.Errorf("\n未通过\n export\t%v\n got\t%v\n", str2, str1)
	}
	fmt.Println("")
}

// 前序遍历树（空节点为null）
func getTreeArr1(root *TreeNode) []string {
	arr := make([]string, 0)
	if root == nil {
		arr = append(arr, "null")
		return arr
	}
	arr = append(arr, strconv.Itoa(root.Val))
	if !(root.Left == nil && root.Right == nil) {
		arr = append(arr, getTreeArr(root.Left)...)
		arr = append(arr, getTreeArr(root.Right)...)
	}
	return arr
}

func getTreeArr(root *TreeNode) []string {
	var ans []string
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		tmp := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		if tmp == nil {
			ans = append(ans, "null")
			continue
		} else {
			ans = append(ans, fmt.Sprintf("%d", tmp.Val))
		}
		if tmp.Left == nil && tmp.Right == nil {
			continue
		}
		queue.PushBack(tmp.Left)
		queue.PushBack(tmp.Right)
	}
	return ans
}

func deserialize(data string) *TreeNode {
	if data == "null" || data == "[]" {
		return nil
	}

	trimmedData := strings.Trim(data, "[]")
	values := strings.Split(trimmedData, ",")

	return buildTree(&values)
}

func buildTree(values *[]string) *TreeNode {
	var layer = make([]*TreeNode, 0)
	var layer2 = make([]*TreeNode, 0)
	first := (*values)[0]
	remain := (*values)[1:]

	val, _ := strconv.Atoi(first)
	node := &TreeNode{Val: val}
	layer = append(layer, node)

	for true {
		if len(remain) <= 0 {
			break
		}
		for true {
			if len(layer) <= 0 {
				break
			}
			var left = remain[0]
			var right = remain[1]
			remain = remain[2:]
			var parent = layer[0]
			layer = layer[1:]
			if left != "null" {
				val, _ = strconv.Atoi(left)
				parent.Left = &TreeNode{Val: val}
				layer2 = append(layer2, parent.Left)
			}
			if right != "null" {
				val, _ = strconv.Atoi(right)
				parent.Right = &TreeNode{Val: val}
				layer2 = append(layer2, parent.Right)
			}
		}
		var tmp = layer
		layer = layer2
		layer2 = tmp
	}
	return node
}

func buildTree1(values *[]string) *TreeNode {
	if len(*values) == 0 {
		return nil
	}

	valStr := (*values)[0]
	*values = (*values)[1:]

	if valStr == "null" {
		return nil
	}

	val, _ := strconv.Atoi(valStr)
	node := &TreeNode{Val: val}
	node.Left = buildTree(values)
	node.Right = buildTree(values)

	return node
}

// func TestDeleteNode(t *testing.T) {
// 	input := "[3,9,20,11,null,15,7,null,null,2,4,6,8]"
// 	root := deserialize(input)
// 	fmt.Println(getTreeArr1(root))
// }

func TestDeleteNode(t *testing.T) {
	tree1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  9,
			Left: &TreeNode{Val: 11},
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   7,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 8},
			},
		},
	}
	tree1_compare := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	tree1 = DeleteNode(tree1)
	equal(t, "树1", getTreeArr1(tree1), getTreeArr1(tree1_compare))

	tree2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 7},
		},
	}
	tree2_compare := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	tree2 = DeleteNode(tree2)
	equal(t, "树2", getTreeArr1(tree2), getTreeArr1(tree2_compare))

	tree3 := &TreeNode{Val: 5}
	var tree3_compare *TreeNode
	tree3 = DeleteNode(tree3)
	equal(t, "树3", getTreeArr1(tree3), getTreeArr1(tree3_compare))

	tree4 := &TreeNode{}
	var tree4_compare *TreeNode
	tree4 = DeleteNode(tree4)
	equal(t, "树4", getTreeArr1(tree4), getTreeArr1(tree4_compare))
}
