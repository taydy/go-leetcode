package medium

import "github.com/taydy/go-leetcode/structure"

func RemoveNthFromEnd(head *structure.ListNode, n int) *structure.ListNode {
	root, cache := head, head
	for ; n > 0; n-- {
		root = root.Next
	}
	if root == nil {
		return cache.Next
	}

	for root.Next != nil {
		root = root.Next
		cache = cache.Next
	}

	cache.Next = cache.Next.Next

	return head
}
