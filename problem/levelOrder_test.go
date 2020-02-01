package problem

func levelOrder(root *TreeNode) [][]int {
	result := make(map[int][]int, 0)
	if root == nil {
		return [][]int{}
	}
	zero := []int{root.Val}
	result[0] = zero
	if root.Left != nil {
		recursiveLevel(root.Left, result, 1)
	}
	if root.Right != nil {
		recursiveLevel(root.Right, result, 1)

	}
	last := make([][]int, len(result))
	for i := 0; i < len(result); i++ {
		last[i] = result[i]
	}
	return last
}

func recursiveLevel(node *TreeNode, result map[int][]int, level int) {
	if node != nil {
		_, ok := result[level]
		if ok {
			result[level] = append(result[level], node.Val)
		} else {
			result[level] = []int{node.Val}
		}
	}
	if node.Left != nil {
		recursiveLevel(node.Left, result, level+1)
	}
	if node.Right != nil {
		recursiveLevel(node.Right, result, level+1)
	}
}
