package problem

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	current := head
	for current != nil {
		nextTemp := current.Next
		current.Next = pre
		pre = current
		current = nextTemp
	}
	return pre
}
