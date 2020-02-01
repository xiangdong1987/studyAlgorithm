package problem

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool)
	var result *ListNode
	if headB == nil || headA == nil {
		return nil
	}
	for headA != nil {
		nodeMap[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		_, ok := nodeMap[headB]
		if ok {
			result = headB
			break
		}
		headB = headB.Next
	}
	return result
}
