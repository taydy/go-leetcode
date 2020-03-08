package medium

import (
	"github.com/taydy/go-leetcode/util"
	"math"
)

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
	//return coinChange1(coins, amount, make([]int, amount))
	return coinChange2(coins, amount)
}

// 动态规划-自上而下
func coinChange1(coins []int, amount int, count []int) int {
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
		r := coinChange1(coins, amount-coin, count)
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

// 动态规划-自下而上
func coinChange2(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	for i := range dp {
		dp[i] = max
	}
	dp[0] = 0

	for i := 0; i < max; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = util.MinInt(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}

}
