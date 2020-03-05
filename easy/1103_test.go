package easy

import (
	"github.com/taydy/go-leetcode/medium"
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	candies := []int{
		7, 10,
	}
	numPeople := []int{
		4, 3,
	}
	excepts := [][]int{
		{1, 2, 3, 1},
		{5, 2, 3},
	}
	for i, v := range candies {
		if result := DistributeCandies(v, numPeople[i]); !medium.Equal(result, excepts[i]) {
			t.Logf("candies %d, num people %d, except %v, but got wrong answer %v", v, numPeople[i], excepts[i], result)
			t.FailNow()
		}
	}
}
