package easy

import "testing"

func TestGcdOfStrings(t *testing.T) {
	str1 := []string{
		"ABCABC",
		"ABABAB",
		"LEET",
	}
	str2 := []string{
		"ABC",
		"ABAB",
		"CODE",
	}
	excepts := []string{
		"ABC",
		"AB",
		"",
	}

	for i := range str1 {
		if result := GcdOfStrings(str1[i], str2[i]); result != excepts[i] {
			t.Logf("str1 %s, str2 %s, except %s, but got wrong answer %s", str1[i], str2[i], excepts[i], result)
			t.FailNow()
		}
	}
}
