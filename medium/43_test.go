package medium

import "testing"

func TestStringMultiply(t *testing.T) {
	result := StringMultiply("2", "3")
	if result != "6" {
		t.Logf("except %s, actual %s", "6", result)
		t.FailNow()
	}

	result = StringMultiply("123", "456")
	if result != "56088" {
		t.Logf("except %s, actual %s", "56088", result)
		t.FailNow()
	}

	result = StringMultiply("0", "0")
	if result != "0" {
		t.Logf("except %s, actual %s", "0", result)
		t.FailNow()
	}
}
