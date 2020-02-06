package problem

import "math"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dsp(root, 1)
}

func dsp(node *TreeNode, height int) int {
	if node == nil {
		return height
	}
	if node.Left == nil && node.Right == nil {
		return height
	}
	return int(math.Max(float64(dsp(node.Left, height+1)), float64(dsp(node.Right, height+1))))
}
