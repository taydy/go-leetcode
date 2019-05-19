package medium

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	str := "babab"
	if result := longestPalindrome(str); result != "bab" {
		fmt.Printf("input %s, expect %s, actual %s \n", str, "bab", result)
		t.FailNow()
	}

	str = "cbbd"
	if result := longestPalindrome(str); result != "bb" {
		fmt.Printf("input %s, expect %s, actual %s \n", str, "bb", result)
		t.FailNow()
	}
}
