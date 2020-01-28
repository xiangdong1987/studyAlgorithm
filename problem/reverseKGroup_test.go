package problem

import (
	"fmt"
	"testing"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	fmt.Println(head)
	current := head
	count := 0
	for count != k && current != nil {
		current = current.Next
		count++
	}
	if count == k {
		current = reverseKGroup(current, k)
		for count > 0 {
			nextTemp := head.Next
			head.Next = current
			current = head
			head = nextTemp
			count--
		}
		head = current
	}

	return head
}
func TestReverseKGroup(t *testing.T) {
	head := &ListNode{}
	head.Val = 1
	Next1 := &ListNode{}
	Next1.Val = 2
	head.Next = Next1
	Next2 := &ListNode{}
	Next2.Val = 3
	Next1.Next = Next2
	Next3 := &ListNode{}
	Next3.Val = 4
	Next2.Next = Next3
	Next4 := &ListNode{}
	Next4.Val = 5
	Next3.Next = Next4
	reverseKGroup(head, 2)
}
