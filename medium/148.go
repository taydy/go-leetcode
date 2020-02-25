package medium

import "github.com/taydy/go-leetcode/structure"

//	在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
//
//	示例 1:
//		输入: 4->2->1->3
//		输出: 1->2->3->4
//
//	示例 2:
//		输入: -1->5->3->4->0
//		输出: -1->0->3->4->5
//
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/sort-list
//
func LinkedSortList(head *structure.ListNode) *structure.ListNode {
	result := &structure.ListNode{
		Val:  -1,
		Next: head,
	}
	length := 0
	temp := head
	for temp != nil {
		temp = temp.Next
		length++
	}

	for size := 1; size < length; size <<= 1 {
		cur := result.Next
		tail := result

		for cur != nil {
			left := cur
			right := cut(left, size)
			cur = cut(right, size)

			tail.Next = merge(left, right)

			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}

	return result.Next
}

func cut(root *structure.ListNode, n int) *structure.ListNode {
	p := root
	for n > 1 && p != nil {
		p = p.Next
		n--
	}

	if p == nil {
		return nil
	}

	next := p.Next
	p.Next = nil
	return next
}

func merge(l1 *structure.ListNode, l2 *structure.ListNode) *structure.ListNode {
	root := &structure.ListNode{
		Val: -1,
	}

	p := root
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			p = p.Next
			l1 = l1.Next
		} else {
			p.Next = l2
			p = p.Next
			l2 = l2.Next
		}
	}

	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}

	return root.Next
}
