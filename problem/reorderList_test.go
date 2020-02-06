package problem

import (
	"testing"
)

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//倒序
	left := head
	right := slow.Next
	slow.Next = nil
	var preNode *ListNode
	current := right
	for current != nil {
		tmpNext := current.Next
		current.Next = preNode
		preNode = current
		current = tmpNext
	}
	right = preNode
	//拼接
	result := &ListNode{}
	for left != nil || right != nil {
		if left != nil {
			result.Next = left
			left = left.Next
			result = result.Next
		}
		if right != nil {
			result.Next = right
			right = right.Next
			result = result.Next
		}
	}
}
func TestReorderList(t *testing.T) {
	node1 := &ListNode{}
	node1.Val = 1
	node2 := &ListNode{}
	node2.Val = 2
	node1.Next = node2
	node3 := &ListNode{}
	node3.Val = 3
	node2.Next = node3
	node4 := &ListNode{}
	node4.Val = 4
	node3.Next = node4
	reorderList(node1)
}
