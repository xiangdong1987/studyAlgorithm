package problem

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var arr []*ListNode
	current := head
	//遍历链表
	for current != nil {
		if len(arr) < n+1 {
			arr = append(arr, current)
		} else {
			arr = append(arr, current)
			arr = arr[1:]
		}
		current = current.Next
	}
	if len(arr) == n {
		head = head.Next
		return head
	}

	//重组链表
	pre := arr[0]
	if n == 1 {
		pre.Next = nil
		return head
	}
	next := arr[2]
	pre.Next = next
	return head
}
