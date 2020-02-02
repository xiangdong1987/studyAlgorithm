package problem

import (
	"fmt"
)

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make(map[int]int, 0)
	result[0] = root.Val
	if root.Right != nil {
		RecursiveRightSide(root.Right, 1, result)
	}
	if root.Left != nil {
		RecursiveRightSide(root.Left, 1, result)
	}
	last := make([]int, len(result))
	for i := 0; i < len(result); i++ {
		last[i] = result[i]
	}
	return last
}
func RecursiveRightSide(node *TreeNode, level int, result map[int]int) {
	_, ok := result[level]
	if !ok {
		fmt.Println(node)
		result[level] = node.Val
	}
	if node.Right != nil {
		RecursiveRightSide(node.Right, level+1, result)
	}
	if node.Left != nil {
		RecursiveRightSide(node.Left, level+1, result)
	}

}
