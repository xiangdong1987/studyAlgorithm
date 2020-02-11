package problem

func sortList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	quickSort(dummy.Next, nil)
	return dummy.Next
}

func quickSort(head *ListNode, end *ListNode) {
	if head.Next == nil || head.Next == end {
		return
	}
	dummy := &ListNode{}
	dummy.Next = head
	min := head
	preNode := dummy
	current := head
	for current != nil {
		if current.Val < min.Val {
			//交换中间值和小于中间值的值
			preNode.Next = current
			current = min
			min.Next = current.Next
		}
		current = current.Next
		if current == end {
			break
		}
	}
	quickSort(min, nil)
	quickSort(dummy.Next, min)
}
