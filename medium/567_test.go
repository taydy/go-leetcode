package medium

import "testing"

func TestCheckInclusion(t *testing.T) {
	if !CheckInclusion("ab", "eidbaooo") {
		t.Logf("except true, actual false")
		t.FailNow()
	}

	if CheckInclusion("ab", "eidboaoo") {
		t.Logf("except false, actual true")
		t.FailNow()
	}

	if !CheckInclusion("adc", "dcda") {
		t.Logf("except true, actual false")
		t.FailNow()
	}
}
