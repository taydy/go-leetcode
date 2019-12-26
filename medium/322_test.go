package medium

import "testing"

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 100
	t.Log(CoinChange(coins, amount))
}
