package problem

/**
112 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
说明: 叶子节点是指没有子节点的节点。
问题分析：
1.看到二叉树，首先第一反应就是递归，因为用递归可以很简单的实现深度遍历
2.也可以用迭代，但是需要额外的空间来做父节点存储，并且最好是栈的数据结构，快速出栈入栈
3.我就使用了迭代的方式进行
*/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		if stack[0].Left != nil {
			stack[0].Left.Val = stack[0].Left.Val + stack[0].Val
			stack = append(stack, stack[0].Left)
		}
		if stack[0].Right != nil {
			stack[0].Right.Val = stack[0].Right.Val + stack[0].Val
			stack = append(stack, stack[0].Right)
		}
		if stack[0].Left == nil && stack[0].Right == nil && stack[0].Val == sum {
			return true
		}

		stack = stack[1:]
	}
	return false
}
