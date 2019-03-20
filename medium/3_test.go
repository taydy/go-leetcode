package medium

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T)  {

	str := "tmmzuxt"
	if result := lengthOfLongestSubstring(str); result != 5 {
		fmt.Printf("input %s, expect %d, actual %d \n", str, 3, result)
		t.FailNow()
	}

	str = "abcabcbb"
	if result := lengthOfLongestSubstring(str); result != 3 {
		fmt.Printf("input %s, expect %d, actual %d \n", str, 3, result)
		t.FailNow()
	}

	str = "bbbbb"
	if result := lengthOfLongestSubstring(str); result != 1 {
		fmt.Printf("input %s, expect %d, actual %d \n", str, 3, result)
		t.FailNow()
	}

	str = "pwwkew"
	if result := lengthOfLongestSubstring(str); result != 3 {
		fmt.Printf("input %s, expect %d, actual %d \n", str, 3, result)
		t.FailNow()
	}

	str = " "
	if result := lengthOfLongestSubstring(str); result != 1 {
		fmt.Printf("input %s, expect %d, actual %d \n", str, 3, result)
		t.FailNow()
	}

	str = "dvdf"
	if result := lengthOfLongestSubstring(str); result != 3 {
		fmt.Printf("input %s, expect %d, actual %d \n", str, 3, result)
		t.FailNow()
	}

	str = "abba"
	if result := lengthOfLongestSubstring(str); result != 2 {
		fmt.Printf("input %s, expect %d, actual %d \n", str, 3, result)
		t.FailNow()
	}
}