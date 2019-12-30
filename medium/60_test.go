package medium

import "testing"

func TestGetPermutation(t *testing.T) {
	result := GetPermutation(3, 3)
	if result != "213" {
		t.Logf("except %s, actual %s", "213", result)
	}

	result = GetPermutation(4, 9)
	if result != "2314" {
		t.Logf("except %s, actual %s", "2314", result)
	}

	result = GetPermutation(3, 5)
	if result != "312" {
		t.Logf("except %s, actual %s", "312", result)
	}
}
