package problem

func inorderTraversal(root *TreeNode) []int {
	return recursiveInorder(root)
}

func recursiveInorder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	var result []int
	if node.Left != nil {
		result = append(result, recursiveInorder(node.Left)...)
	}
	result = append(result, node.Val)
	if node.Right != nil {
		result = append(result, recursiveInorder(node.Right)...)
	}
	return result
}
