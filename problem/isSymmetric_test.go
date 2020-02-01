package problem

func isSymmetric(root *TreeNode) bool {
	return recursiveSymmetric(root.Left, root.Right)
}

func recursiveSymmetric(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 != nil && node2 != nil {
		if recursiveSymmetric(node1.Left, node2.Right) && recursiveSymmetric(node1.Right, node2.Left) && node1.Val == node2.Val {
			return true
		}
	} else {
		if node1 == node2 {
			return true
		}
	}
	return false
}
