package medium

import "taydy/go-leetcode/util"

//	给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列。
//
//	一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//	例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。
//
//	若这两个字符串没有公共子序列，则返回 0。
//
//	示例 1:
//		输入：text1 = "abcde", text2 = "ace"
//		输出：3
//		解释：最长公共子序列是 "ace"，它的长度为 3。
//	示例 2:
//		输入：text1 = "abc", text2 = "abc"
//		输出：3
//		解释：最长公共子序列是 "abc"，它的长度为 3。
//	示例 3:
//		输入：text1 = "abc", text2 = "def"
//		输出：0
//		解释：两个字符串没有公共子序列，返回 0。
//
//	提示:
//		1 <= text1.length <= 1000
//		1 <= text2.length <= 1000
//		输入的字符串只含有小写英文字符。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/longest-common-subsequence
//
func LongestCommonSubsequence(text1 string, text2 string) int {

	c1 := []rune(text1)
	c2 := []rune(text2)

	dp := make([][]int, 0)
	for i := 0; i < len(c1)+1; i++ {
		dp = append(dp, make([]int, len(c2)+1))
	}

	for r := 1; r < len(c1)+1; r++ {
		for c := 1; c < len(c2)+1; c++ {
			if c1[r-1] == c2[c-1] {
				dp[r][c] = dp[r-1][c-1] + 1
			} else {
				dp[r][c] = util.MaxInt(dp[r-1][c], dp[r][c-1])
			}
		}
	}

	return dp[len(c1)][len(c2)]
}