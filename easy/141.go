package easy

import "github.com/taydy/go-leetcode/structure"

//	给定一个链表，判断链表中是否有环。
//
//	为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/linked-list-cycle
//
func hasCycle(head *structure.ListNode) bool {
	if head == nil {
		return false
	}
	quick, slow := head, head
	for quick.Next != nil && quick.Next.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next

		if quick == slow {
			return true
		}
	}
	return false
}
