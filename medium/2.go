package medium

//
// 给出两个非空的链表用来表示两个非负的整数。其中，它们各自的位数是按照逆序的方式存储的，并且它们的每个节点只能存储一位数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
//	 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//	 输出：7 -> 0 -> 8
//	 原因：342 + 465 = 807
//


//
// 复杂度分析
//
//  时间复杂度：O(max(m, n))，假设 m 和 n 分别表示 l1 和 l2 的长度，下面的算法最多重复 max(m, n) 次。
//
//  空间复杂度：O(max(m, n)))， 新列表的长度最多为 max(m,n) + 1。
//
//
//

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0
	for ; l1 != nil || l2 != nil; {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		carry = sum / 10
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
	}
	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}