package medium

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	l1 := []int{2, 4, 3}
	l2 := []int{5, 6, 4}
	result := addTwoNumbers(arrayToList(l1), arrayToList(l2))
	if !equal(ListToArray(result), []int{7, 0, 8}) {
		t.FailNow()
	}

	l1 = []int{1, 8}
	l2 = []int{}
	result = addTwoNumbers(arrayToList(l1), arrayToList(l2))
	if !equal(ListToArray(result), []int{1, 8}) {
		t.FailNow()
	}

	l1 = []int{5}
	l2 = []int{5}
	result = addTwoNumbers(arrayToList(l1), arrayToList(l2))
	if !equal(ListToArray(result), []int{0, 1}) {
		t.FailNow()
	}
}

func arrayToList(arr []int) *ListNode {
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

func equal(x, y []int) bool {
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