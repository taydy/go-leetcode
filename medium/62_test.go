package medium

import "testing"

func TestUniquePaths(t *testing.T) {
	m, n := 3, 2
	except := 3
	if result := UniquePaths(m, n); result != except {
		t.Logf("except %d, actual %d", except, result)
		t.FailNow()
	}

	m, n = 7, 3
	except = 28
	if result := UniquePaths(m, n); result != except {
		t.Logf("except %d, actual %d", except, result)
		t.FailNow()
	}
}
