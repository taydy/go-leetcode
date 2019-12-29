package easy

import "testing"

func TestAddStrings(t *testing.T) {
	result := AddStrings("123", "27")
	if result != "150" {
		t.Logf("except 150, actual %s", result)
		t.FailNow()
	}

	result = AddStrings("1", "9")
	if result != "10" {
		t.Logf("except 10, actual %s", result)
		t.FailNow()
	}
}
