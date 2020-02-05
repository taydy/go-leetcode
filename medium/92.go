package medium

//	反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
//	说明:
//		1 ≤ m ≤ n ≤ 链表长度。
//
//	示例:
//
//		输入: 1->2->3->4->5->NULL, m = 2, n = 4
//		输出: 1->4->3->2->5->NULL
//
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
//
func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	cur := head

	for m > 1 {
		prev = cur
		cur = cur.Next
		m--
		n--
	}

	con, tail := prev, cur

	for ; cur != nil && n > 0; n-- {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	if con != nil {
		con.Next = prev
	} else {
		head = prev
	}

	tail.Next = cur

	return head
}
