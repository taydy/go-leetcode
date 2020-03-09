package hard

import (
	"github.com/taydy/go-leetcode/util"
	"math"
)

func MaxProfitIII(prices []int) int {
	n := len(prices)

	if n == 0 {
		return 0
	}

	k := 2
	dp := make([][][]int, 0)
	for i := 0; i < n; i++ {
		tmp := make([][]int, 0)
		for j := 0; j <= k; j++ {
			tmp = append(tmp, make([]int, 2))
		}
		dp = append(dp, tmp)
	}

	dp[0][1][0], dp[0][1][1] = 0, -prices[0]
	dp[0][2][0], dp[0][2][1] = 0, math.MinInt32

	for i := 1; i < n; i++ {
		dp[i][1][0] = util.MaxInt(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
		dp[i][1][1] = util.MaxInt(dp[i-1][1][1], -prices[i])
		dp[i][2][0] = util.MaxInt(dp[i-1][2][0], dp[i-1][2][1]+prices[i])
		dp[i][2][1] = util.MaxInt(dp[i-1][2][1], dp[i-1][1][0]-prices[i])
	}

	return util.MaxInt(dp[n-1][2][0], dp[n-1][1][0])
}
