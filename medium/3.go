package medium

//
//	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//	示例 1:
//
//		输入: "abcabcbb"
//		输出: 3
//		解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
// 	示例 2:
//
//		输入: "bbbbb"
//		输出: 1
//		解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
// 	示例 3:
//
//		输入: "pwwkew"
//		输出: 3
//		解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//

func lengthOfLongestSubstring(s string) int {

	length := len(s)
	chars := []rune(s)
	max := 0
	index := make([]int, 128, 128)
	for start, end := 0, 0; end < length; end++ {
		if value := index[chars[end]]; value > start {
			start = value
		}
		if temp := end - start + 1; temp > max {
			max = temp
		}
		index[chars[end]] = end + 1
	}

	return max
}
