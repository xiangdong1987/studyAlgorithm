package problem

func zigzagLevelOrder(root *TreeNode) [][]int {
	nextLevel := []*TreeNode{root}
	var result [][]int
	level := 0
	for len(nextLevel) > 0 {
		var next []*TreeNode
		var tmp []int
		for i := 0; i < len(nextLevel); i++ {
			//从左到右
			if level%2 == 0 {
				if level == 0 {
					if nextLevel[0] == nil {
						return result
					}
					if nextLevel[0].Right != nil {
						next = append(next, nextLevel[0].Right)
					}
					if nextLevel[0].Left != nil {
						next = append(next, nextLevel[0].Left)
					}
					tmp = append(tmp, nextLevel[0].Val)
					break
				}
				if nextLevel[i].Left != nil {
					next = append([]*TreeNode{nextLevel[i].Left}, next...)
				}
				if nextLevel[i].Right != nil {
					next = append([]*TreeNode{nextLevel[i].Right}, next...)
				}
			} else {
				if nextLevel[i].Right != nil {
					next = append([]*TreeNode{nextLevel[i].Right}, next...)
				}
				if nextLevel[i].Left != nil {
					next = append([]*TreeNode{nextLevel[i].Left}, next...)
				}
			}
			tmp = append(tmp, nextLevel[i].Val)
		}
		nextLevel = next
		result = append(result, tmp)
		level++
	}
	return result
}
