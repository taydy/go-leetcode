package medium

import (
	"strconv"
	"strings"
)

func GetPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	num := 0
	temp := permuteBacktrack1(nums, 0, n, &num, k)

	var r strings.Builder
	for _, v := range temp {
		r.WriteString(strconv.Itoa(v))
	}

	return r.String()
}

func permuteBacktrack1(nums []int, first int, length int, n *int, k int) []int {
	if first == length {
		*n++
		if *n == k {
			return nums
		}
	}

	for i := first; i < length; i++ {
		nums[first], nums[i] = nums[i], nums[first]
		if temp := permuteBacktrack1(nums, first+1, length, n, k); temp != nil {
			return temp
		}
		nums[first], nums[i] = nums[i], nums[first]
	}
	return nil
}
