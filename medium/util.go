package medium

import "github.com/taydy/go-leetcode/structure"

func ArrayToList(arr []int) *structure.ListNode {
	result := &structure.ListNode{}
	p := result
	for _, value := range arr {
		if p == nil {
			p = &structure.ListNode{Val: value}
		} else {
			p.Next = &structure.ListNode{Val: value}
			p = p.Next
		}
	}
	return result.Next
}

func ListToArray(l1 *structure.ListNode) []int {
	result := make([]int, 0)
	for ; l1 != nil; l1 = l1.Next {
		result = append(result, l1.Val)
	}
	return result
}

func Equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
