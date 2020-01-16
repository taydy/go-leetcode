package easy

//	编写一个程序，找到两个单链表相交的起始节点。
func GetIntersectionNode(headA, headB *ListNode) *ListNode {

	nodeA, nodeB := headA, headB
	for nodeA != nodeB {

		if nodeA != nil {
			nodeA = nodeA.Next
		} else {
			nodeA = headB
		}

		if nodeB != nil {
			nodeB = nodeB.Next
		} else {
			nodeB = headA
		}
	}

	return nodeA

}
