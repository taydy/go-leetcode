package easy

//	反转一个单链表。
//
//	示例:
//
//		输入: 1->2->3->4->5->NULL
//		输出: 5->4->3->2->1->NULL

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {

	// 前指针节点
	var prev *ListNode
	// 当前指针节点
	cur := head

	for cur != nil {
		// 临时节点，暂存当前节点的下一节点，用于后移
		temp := cur.Next
		// 将当前节点指向它前面的节点
		cur.Next = prev
		// 前指针后移
		prev = cur
		// 当前指针节点后移
		cur = temp
	}

	return prev
}
