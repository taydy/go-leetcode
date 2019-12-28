package medium

import "testing"

func TestGenerateMatrix(t *testing.T) {
	n := 3
	t.Logf("%+v", GenerateMatrix(n))
}
