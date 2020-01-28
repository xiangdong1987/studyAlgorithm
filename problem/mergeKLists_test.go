package problem

import "fmt"

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	pre := head
	var min *ListNode
	var pos int
	for len(lists) > 0 {
		if len(lists) == 1 {
			fmt.Println(pre, lists[0])
			if lists[0] == nil {
				break
			}
			pre.Next = lists[0]
			break
		}
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				lists = append(lists[:i], lists[i+1:]...)
				break
			}
			if i == 0 {
				min = lists[0]
				pos = i
			}
			if lists[i].Val <= min.Val {
				min = lists[i]
				pos = i
			}
			if i == len(lists)-1 {
				pre.Next = min
				if min.Next != nil {
					lists[pos] = lists[pos].Next
				} else {
					lists = append(lists[:pos], lists[pos+1:]...)
				}
				pre = pre.Next
			}
		}
	}
	return head.Next
}
