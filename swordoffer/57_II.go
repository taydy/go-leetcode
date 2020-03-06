package swordoffer

//	输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
//	序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
//
//	示例 1：
//		输入：target = 9
//		输出：[[2,3,4],[4,5]]
//	示例 2：
//		输入：target = 15
//		输出：[[1,2,3,4,5],[4,5,6],[7,8]]
//
//	限制：
//		1 <= target <= 10^5
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
//
func FindContinuousSequence(target int) [][]int {
	results := make([][]int, 0)
	var ans []int

	for l, r := 1, 2; l < r; {
		sum := ((l + r) * (r - l + 1)) >> 1
		if sum == target {
			ans = make([]int, 0)
			for i := l; i <= r; i++ {
				ans = append(ans, i)
			}
			results = append(results, ans)
			l++
		} else if sum < target {
			r++
		} else {
			l++
		}
	}

	return results
}
