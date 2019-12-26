package medium

import "math"

//	给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//
//	示例 1:
//		输入: coins = [1, 2, 5], amount = 11
//		输出: 3
//		解释: 11 = 5 + 5 + 1
//
//	示例 2:
//		输入: coins = [2], amount = 3
//		输出: -1
//
//	说明:
//		你可以认为每种硬币的数量是无限的。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/coin-change
//
func CoinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	return coinChange(coins, amount, make([]int, amount))
}

func coinChange(coins []int, amount int, count []int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}

	if c := count[amount-1]; c != 0 {
		return c
	}

	min := math.MaxInt32
	for _, coin := range coins {
		r := coinChange(coins, amount-coin, count)
		if r != -1 && r < min {
			min = r + 1
		}
	}

	if min == math.MaxInt32 {
		min = -1
	}

	count[amount-1] = min
	return min
}
