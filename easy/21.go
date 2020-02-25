package easy

import "github.com/taydy/go-leetcode/structure"

// 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//示例：
//
//	输入：1->2->4, 1->3->4
//	输出：1->1->2->3->4->4

func MergeTwoLists(l1 *structure.ListNode, l2 *structure.ListNode) *structure.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var result *structure.ListNode
	if l1.Val <= l2.Val {
		result = l1
		result.Next = MergeTwoLists(l1.Next, l2)
	} else {
		result = l2
		result.Next = MergeTwoLists(l1, l2.Next)
	}
	return result
}
