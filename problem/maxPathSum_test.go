package problem

import "math"

var ans int

/**
124 给定一个非空二叉树，返回其最大路径和。
本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

问题分析：
1.这道题就是一个经典的深度递归问题，树肯定是要递归了，深度或者广度
2.因为我们要求的就是最大路径，所以我们肯定要进行深度的遍历
3.我总结了一些我关于深度递归问题的解法，总结为一句话
	思路要正向，执行要逆向
4.思路正向就是说，我们假设递归内的递归经完成，在这个例子里，我们就假设左右节点返回的就是
他们的最优解，我们的情况就分为了两种：
	经过当前节点的左右节点
	经过当前节点的左右节点中较大的一个
这样我们就覆盖了所有的情况，递归方法如下代码
*/
func maxPathSum(root *TreeNode) int {
	ans = INT32_MIN
	dfsMax(root)
	return ans
}

func dfsMax(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := int(math.Max(float64(dfsMax(node.Left)), 0))
	right := int(math.Max(float64(dfsMax(node.Right)), 0))
	newPath := node.Val + left + right
	ans = int(math.Max(float64(newPath), float64(ans)))
	return node.Val + int(math.Max(float64(left), float64(right)))
}
