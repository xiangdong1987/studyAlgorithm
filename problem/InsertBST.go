package problem

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	myRoot := root
	RecursiveInsert(root, val)
	return myRoot
}

func RecursiveInsert(root *TreeNode, val int) {
	fmt.Println(root)
	if val < root.Val {
		if root.Left != nil {
			insertIntoBST(root.Left, val)
		} else {
			newNode := TreeNode{}
			newNode.Val = val
			root.Left = &newNode
		}
	} else {
		if root.Right != nil {
			insertIntoBST(root.Right, val)
		} else {
			newNode := TreeNode{}
			newNode.Val = val
			root.Right = &newNode
		}
	}
}
