package medium

import "github.com/taydy/go-leetcode/structure"

// 给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。
//
//	你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//	进阶:
//		如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
//	示例:
//		输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
//		输出: 7 -> 8 -> 0 -> 7
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/add-two-numbers-ii
//
// 思路:
// 	1. 将短的链表前面补零, 让两个链表等长
//	2. 递归求解
func AddTwoNumbers2(l1 *structure.ListNode, l2 *structure.ListNode) *structure.ListNode {
	size1 := getListNodeSize(l1)
	size2 := getListNodeSize(l2)

	if size1 > size2 {
		l2 = fillZero(l2, size1-size2)
	} else if size1 < size2 {
		l1 = fillZero(l1, size2-size1)
	}
	head := addTwoNumber(l1, l2, &structure.ListNode{
		Val: 0,
	})
	if head.Val == 0 {
		head = head.Next
	}
	return head
}

func addTwoNumber(l1 *structure.ListNode, l2 *structure.ListNode, head *structure.ListNode) *structure.ListNode {
	if l1 == nil || l2 == nil {
		head = &structure.ListNode{
			Val: 0,
		}
		return head
	}

	head = addTwoNumber(l1.Next, l2.Next, head)
	val := l1.Val + l2.Val + head.Val
	head.Val = val % 10
	head = &structure.ListNode{
		Val:  val / 10,
		Next: head,
	}
	return head
}

// 在 list 前面补零
func fillZero(head *structure.ListNode, num int) *structure.ListNode {
	for ; num > 0; num-- {
		head = &structure.ListNode{
			Val:  0,
			Next: head,
		}
	}
	return head
}

// 获取 list 长度
func getListNodeSize(head *structure.ListNode) int {
	size := 0
	node := head
	for node != nil {
		node = node.Next
		size++
	}
	return size
}
