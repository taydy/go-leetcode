package medium

func ArrayToList(arr []int) *ListNode {
	result := &ListNode{}
	p := result
	for _, value := range arr {
		if p == nil {
			p = &ListNode{Val: value}
		} else {
			p.Next = &ListNode{Val: value}
			p = p.Next
		}
	}
	return result.Next
}

func ListToArray(l1 *ListNode) []int {
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
