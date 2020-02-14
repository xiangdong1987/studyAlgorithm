package problem

/**
92 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
问题分析：
1.首先看到的是单向链表相关，单向链表最难的是边界问题，主要问题是要知道先前节点和终点，所以为了解决这个问题
我们一般会设置哨兵节点，这样的话自己虚拟出来的头节点就解决了，第一次的时候没有先前节点的问题，简化程序的写法
2.第二个问题就是链表的指针的交换问题，有时候会交换乱掉，
我自己的理解是，一般要三个指针，先前指针，当前指针和要交换的指针组成三角交换利用Next进行桥接
	pre.Next = current.Next
	tmpHead := leftHead.Next
	leftHead.Next = current
	current.Next = tmpHead
	current = pre.Next
3.这道题还有一个问题就是，什么时候开始交换和什么时候结束交换，我用了两个标记位进行标记
链表翻转
*/

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	dummy := &ListNode{}
	var leftHead *ListNode
	dummy.Next = head
	current := dummy.Next
	pre := dummy
	isTurn := false
	isFinish := false
	num := 1
	for {
		if !isTurn {
			if num == m {
				isTurn = true
				leftHead = pre
			}
			pre = pre.Next
			current = current.Next
		} else {
			if num == n {
				isFinish = true
			}
			pre.Next = current.Next
			tmpHead := leftHead.Next
			leftHead.Next = current
			current.Next = tmpHead
			current = pre.Next
			if isFinish {
				break
			}
		}
		num++
	}
	return dummy.Next
}
