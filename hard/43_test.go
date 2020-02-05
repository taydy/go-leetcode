package hard

import "testing"

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	except := 6
	if result := Trap(height); result != except {
		t.Logf("except %d, actual %d", except, result)
		t.FailNow()
	}

	height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	except = 6
	if result := Trap1(height); result != except {
		t.Logf("except %d, actual %d", except, result)
		t.FailNow()
	}

	height = []int{2, 0, 2}
	except = 2
	if result := Trap(height); result != except {
		t.Logf("except %d, actual %d", except, result)
		t.FailNow()
	}

	height = []int{2, 0, 2}
	except = 2
	if result := Trap1(height); result != except {
		t.Logf("except %d, actual %d", except, result)
		t.FailNow()
	}
}
