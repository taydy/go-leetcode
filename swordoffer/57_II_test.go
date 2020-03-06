package swordoffer

import (
	"testing"
)

func TestFindContinuousSequence(t *testing.T) {
	targets := []int{9, 15}
	for _, v := range targets {
		t.Logf("target %d, output %v", v, FindContinuousSequence(v))
	}

}
