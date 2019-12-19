package medium

import "testing"

func TestMyCircularQueue(t *testing.T) {
	myCircularQueue := Constructor(3)

	if !myCircularQueue.IsEmpty() {
		t.FailNow()
	}

	if !myCircularQueue.EnQueue(1) {
		t.FailNow()
	}
	if !myCircularQueue.EnQueue(2) {
		t.FailNow()
	}
	if !myCircularQueue.EnQueue(3) {
		t.FailNow()
	}
	if myCircularQueue.EnQueue(4) {
		t.FailNow()
	}

	if myCircularQueue.Front() != 1 {
		t.FailNow()
	}

	if myCircularQueue.IsEmpty() {
		t.FailNow()
	}

	if !myCircularQueue.DeQueue() || !myCircularQueue.DeQueue() || !myCircularQueue.DeQueue() {
		t.FailNow()
	}

	if !myCircularQueue.IsEmpty() {
		t.FailNow()
	}

	if myCircularQueue.Front() != -1 {
		t.FailNow()
	}
}
