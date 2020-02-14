package medium

// https://leetcode-cn.com/problems/longest-palindromic-substring/
// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
// 示例 1：
//	 输入: "babad"
//	 输出: "bab"
//	 注意: "aba" 也是一个有效答案。
//
// 示例 2：
//   输入: "cbbd"
//   输出: "bb"
//
//
// 思路:
//	中心扩展算法:
//		回文中心的两侧互为镜像。因此，回文可以从他的中心展开，并且只有2n-1个这样的中心(一个元素为中心的情况有n个，两个元素为中心的情况有n-1个)
//
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	chars := []rune(s)
	var start, end int

	for i := 0; i < len(s); i++ {
		length := expandAroundCenter(chars, i, i)
		lengthB := expandAroundCenter(chars, i, i+1)
		if length < lengthB {
			length = lengthB
		}
		if length > end-start {
			start = i - (length-1)>>1
			end = i + length>>1
		}
	}
	return string(chars[start : end+1])
}

func expandAroundCenter(chars []rune, left, right int) int {
	length := len(chars)
	for left >= 0 && right < length && chars[left] == chars[right] {
		left--
		right++
	}
	return right - left - 1
}
