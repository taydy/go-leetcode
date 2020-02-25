package hard

import "testing"

func TestMaximalRectangle(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	except := 6
	if result := MaximalRectangle(matrix); result != except {
		t.Logf("except %d, actual %d", except, result)
		t.FailNow()
	}

	matrix = [][]byte{
		{'1', '0'},
	}
	except = 1
	if result := MaximalRectangle(matrix); result != except {
		t.Logf("except %d, actual %d", except, result)
		t.FailNow()
	}
}
