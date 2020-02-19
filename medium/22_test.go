package medium

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	n := 3
	t.Logf("n = %d, results = %v", n, GenerateParenthesis(n))

	n = 4
	t.Logf("n = %d, results = %v", n, GenerateParenthesis(n))
}
