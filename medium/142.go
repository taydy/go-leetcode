package medium

//	给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
//	为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//
//	说明：不允许修改给定的链表。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
//
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			fast = head
			for fast != slow {
				fast, slow = fast.Next, slow.Next
			}
			return fast
		}
	}

	return nil
}
