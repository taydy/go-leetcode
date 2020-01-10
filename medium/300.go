package medium

import "taydy/go-leetcode/util"

//	给定一个无序的整数数组，找到其中最长上升子序列的长度。
//
//	示例:
//		输入: [10,9,2,5,3,7,101,18]
//		输出: 4
//		解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
//	说明:
//		可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
//		你算法的时间复杂度应该为 O(n2) 。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
//
// 注: 此题可以使用动态规划+二分查找实现复杂度 O(NlogN) 的解法.
// https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/dong-tai-gui-hua-she-ji-fang-fa-zhi-pai-you-xi-jia/
func LengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	dp := make([]int, n)
	for i, _ := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = util.MaxInt(dp[i], dp[j]+1)
			}
		}
	}

	max := 1
	for _, v := range dp {
		if max < v {
			max = v
		}
	}
	return max
}

func LengthOfLIS1(nums []int) int {
	n := len(nums)
	top := make([]int, n)
	piles := 0
	for i := 0; i < n; i++ {
		poker := nums[i]
		left, right := 0, piles

		for left < right {
			mid := (left + right) >> 1
			if top[mid] >= poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			}
		}

		if left == piles {
			piles++
		}

		top[left] = poker

	}

	return piles
}
