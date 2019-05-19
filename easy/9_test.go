package easy

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	input := 121
	if result := IsPalindrome(input); !result {
		fmt.Printf("input %d, expect %t, actual %t \n", input, true, result)
		t.FailNow()
	}

	input = -121
	if result := IsPalindrome(input); result {
		fmt.Printf("input %d, expect %t, actual %t \n", input, false, result)
		t.FailNow()
	}

	input = 10
	if result := IsPalindrome(input); result {
		fmt.Printf("input %d, expect %t, actual %t \n", input, false, result)
		t.FailNow()
	}

	input = 1
	if result := IsPalindrome(input); !result {
		fmt.Printf("input %d, expect %t, actual %t \n", input, true, result)
		t.FailNow()
	}
}
