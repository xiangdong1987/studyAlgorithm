package problem

import "fmt"

/**
445 给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。
将这两数相加会返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。
进阶:
如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
问题分析：
1.拿到问题第一个想到的就是将两个链表逆序，因为这样就可以从低位向高位进行10进制计算，比较简单
2.但是题目要求不改变原来链表的情况下进行加和这样的话就比较困难
3.我的第一反应思路是，将链表直接转化成int进行加和，可想而知是多么的愚蠢，当遇到int最大限制时就会出错
4.通过这道题，我深刻理解了数据结构对算法的限制，还有链表的灵活运用，可以用链表来进行超大数计算

*/
func addTwoNumbersThree(l1 *ListNode, l2 *ListNode) *ListNode {
	nums1 := make([]int, 0)
	nums2 := make([]int, 0)
	//统计两个链表的长度
	for l1 != nil || l2 != nil {
		if l1 != nil {
			nums1 = append(nums1, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			nums2 = append(nums2, l2.Val)
			l2 = l2.Next
		}
	}
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	lengthN1 := len(nums1)
	lengthN2 := len(nums2)
	pre := 0
	//反向相加
	for i := 0; i < lengthN2; i++ {
		fmt.Println(nums2)
		if lengthN1-1-i >= 0 {
			nums2[lengthN2-1-i] = nums2[lengthN2-1-i] + nums1[lengthN1-1-i] + pre
		} else {
			nums2[lengthN2-1-i] = nums2[lengthN2-1-i] + pre
		}
		if nums2[lengthN2-1-i] >= 10 {
			nums2[lengthN2-1-i] = nums2[lengthN2-1-i] % 10
			pre = 1
		} else {
			pre = 0
		}
	}
	if pre != 0 {
		nums2 = append([]int{1}, nums2...)
	}
	dummy := &ListNode{}
	current := dummy
	for i := 0; i < len(nums2); i++ {
		node := &ListNode{}
		node.Val = nums2[i]
		current.Next = node
		current = current.Next
	}
	return dummy.Next
}

func addTwoNumbersTow(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := 0
	num2 := 0
	//计算链表长度
	for l1 != nil || l2 != nil {
		if l1 != nil {
			num1 = num1*10 + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = num2*10 + l2.Val
			l2 = l2.Next
		}
	}
	sum := num1 + num2
	current := sum
	dummy := &ListNode{}
	if current == 0 {
		dummy.Val = 0
		return dummy
	}
	//数字转链表
	for current != 0 {
		node := &ListNode{}
		node.Val = current % 10
		tmp := dummy.Next
		dummy.Next = node
		node.Next = tmp
		current = current / 10
	}
	return dummy.Next
}
