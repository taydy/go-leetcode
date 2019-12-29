package easy

import "testing"

func TestLongestPalindrome(t *testing.T) {
	s := "abccccdd"
	result := LongestPalindrome(s)
	if result != 7 {
		t.Logf("except %d, actual %d", 7, result)
		t.FailNow()
	}
}
