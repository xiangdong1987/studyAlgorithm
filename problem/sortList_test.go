package problem

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//寻找中点
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var mid *ListNode
	mid, slow.Next = slow.Next, nil
	//切割
	left, right := sortList(head), sortList(mid)
	//合并
	result := &ListNode{}
	current := result
	for left != nil && right != nil {
		if left.Val > right.Val {
			current.Next, right = right, right.Next
		} else {
			current.Next, left = left, left.Next
		}
		current = current.Next
	}
	//将剩下的链接
	if left != nil {
		current.Next = left
	} else {
		current.Next = right
	}
	return result.Next
}
