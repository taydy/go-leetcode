package medium

import "testing"

func TestReverseWords(t *testing.T) {
	result := ReverseWords("the sky is blue")
	actual := "blue is sky the"
	if result != actual {
		t.Logf("except \"%s\", actual \"%s\"", actual, result)
		t.FailNow()
	}

	result = ReverseWords("  hello world!  ")
	actual = "world! hello"
	if result != actual {
		t.Logf("except \"%s\", actual \"%s\"", actual, result)
		t.FailNow()
	}

	result = ReverseWords("a good   example")
	actual = "example good a"
	if result != actual {
		t.Logf("except \"%s\", actual \"%s\"", actual, result)
		t.FailNow()
	}
}
