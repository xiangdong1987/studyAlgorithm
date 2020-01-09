package problem

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	finish := false
	carry := 0
	result := &ListNode{}
	var current *ListNode
	for i := 0; finish != true; i++ {
		tmp := &ListNode{}
		if l1 == nil && l2 == nil {
			finish = true
			if carry == 1 && finish == true {
				carryTmp := &ListNode{}
				carryTmp.Val = 1
				current.Next = carryTmp
			}
			break
		}
		if l1 != nil {
			tmp.Val = tmp.Val + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp.Val = tmp.Val + l2.Val
			l2 = l2.Next
		}
		if carry != 0 {
			tmp.Val = tmp.Val + 1
			carry = 0
		}
		if tmp.Val >= 10 {
			tmp.Val = tmp.Val - 10
			carry = 1
		}
		if i == 0 {
			result = tmp
			current = tmp
		} else {
			current.Next = tmp
			current = tmp
		}
	}
	return result
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{}
	fmt.Print(l1)
}
