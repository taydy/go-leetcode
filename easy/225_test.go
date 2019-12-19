package easy

import "testing"

func TestMyStack(t *testing.T) {
	stack := Constructor()
	if !stack.Empty() {
		t.FailNow()
	}

	stack.Push(1)
	stack.Push(2)

	if stack.Top() != 2 {
		t.FailNow()
	}

	if stack.Pop() != 2 {
		t.FailNow()
	}

	if stack.Pop() != 1 {
		t.FailNow()
	}
}
