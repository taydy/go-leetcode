package swordoffer

import "testing"

func TestMaxQueue(t *testing.T) {
	maxQueue := Constructor()
	maxQueue.Push_back(1)
	maxQueue.Push_back(2)

	if maxQueue.Max_value() != 2 {
		t.FailNow()
	}

	if maxQueue.Pop_front() != 1 {
		t.FailNow()
	}

	if maxQueue.Max_value() != 2 {
		t.FailNow()
	}

	if maxQueue.Pop_front() != 2 {
		t.FailNow()
	}

	if maxQueue.Pop_front() != -1 {
		t.FailNow()
	}

	if maxQueue.Max_value() != -1 {
		t.FailNow()
	}
}
