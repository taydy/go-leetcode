package stack

import "testing"

func TestStack(t *testing.T) {
	stack := New()

	if !stack.IsEmpty() {
		t.Logf("stack should be empty!")
		t.FailNow()
	}

	nums := []int{1, 2, 3, 4, 5, 6}
	for _, v := range nums {
		stack.Push(v)
	}

	if stack.IsEmpty() || stack.Length() != len(nums) {
		t.Logf("stack length is not correct.")
		t.FailNow()
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if stack.Pop() != nums[i] {
			t.FailNow()
		}
	}

}
