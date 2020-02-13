package problem

/**
236 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，
最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大
（一个节点也可以是它自己的祖先）。”
问题分析：
1.一般二叉树节点相关肯定会使用到二叉树的遍历
2.遍历一般分为三种，前序，中序，后序三种遍历方式
3.本体是查找根节点所以先遍历左右节点在遍历父节点 一般采取中序遍历
4.为了能返回上一层，肯定需要记录上一层的节点，所以需要用额外空间来记录，
一般用栈的模式先进后出
我的思路是分别查找两个节点，存储他们的父节点遍历情况，在通过两个栈来查找公共父节点
github
https://github.com/xiangdong1987/studyAlgorithm
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	pStack := recursiveLowest(root, p)
	qStack := recursiveLowest(root, q)
	for i := len(pStack) - 1; i >= 0; i-- {
		for j := len(qStack) - 1; j >= 0; j-- {
			if pStack[i].Val == qStack[j].Val || qStack[j].Val == p.Val {
				return qStack[j]
			}
		}
	}
	return nil
}

func recursiveLowest(root *TreeNode, node *TreeNode) []*TreeNode {
	stack := make([]*TreeNode, 0)
	if root.Val == node.Val {
		stack = append(stack, root)
		return stack
	}
	if root.Left != nil {
		left := recursiveLowest(root.Left, node)
		if len(left) > 0 {
			return append([]*TreeNode{root}, left...)
		}
	}
	if root.Right != nil {
		right := recursiveLowest(root.Right, node)
		if len(right) > 0 {
			return append([]*TreeNode{root}, right...)
		}
	}
	return stack
}
