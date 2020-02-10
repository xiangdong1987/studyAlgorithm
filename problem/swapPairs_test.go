package problem

func swapPairs(head *ListNode) *ListNode {
	//哨兵节点
	dummy := &ListNode{}
	current := head
	dummy.Next = current
	preNode := dummy
	for current != nil && current.Next != nil {
		one := current
		tow := current.Next
		//交换
		preNode.Next = tow
		one.Next = tow.Next
		tow.Next = one
		//下个循环
		preNode = one
		current = one.Next
	}
	return dummy.Next
}
