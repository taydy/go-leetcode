package easy

import "github.com/taydy/go-leetcode/structure"

// 请判断一个链表是否为回文链表。
//
//	示例 1:
//		输入: 1->2
//		输出: false
//	示例 2:
//		输入: 1->2->2->1
//		输出: true

func IsPalindromeList(head *structure.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 快慢指针, 快指针一次走两步,慢指针一次走一步
	fast, slow := head, head

	// 根据快慢指针找到中点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 反转慢指针后面的节点
	slow = ReverseList(slow)

	for slow != nil {
		if head.Val != slow.Val {
			return false
		}
		head, slow = head.Next, slow.Next
	}

	return true
}
