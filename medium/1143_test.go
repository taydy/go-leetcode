package medium

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	text1 := []string{
		"abcde", "abc", "abc",
	}
	text2 := []string{
		"ace", "abc", "def",
	}
	excepts := []int{3, 3, 0}

	for i, _ := range text1 {
		if result := LongestCommonSubsequence(text1[i], text2[i]); result != excepts[i] {
			t.Logf("text1 %s, text2 %s, except %d, actual %d", text1[i], text2[i], excepts[i], result)
			t.FailNow()
		}
	}
}
