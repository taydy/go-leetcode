package medium

//	给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
//	换句话说，第一个字符串的排列之一是第二个字符串的子串。
//
//	示例1:
//		输入: s1 = "ab" s2 = "eidbaooo"
//		输出: True
//		解释: s2 包含 s1 的排列之一 ("ba").
//
//	示例2:
//		输入: s1= "ab" s2 = "eidboaoo"
//		输出: False
//
//	注意：
//		输入的字符串只包含小写字母
//		两个字符串的长度都在 [1, 10,000] 之间
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/permutation-in-string
//
func CheckInclusion(s1 string, s2 string) bool {
	length1 := len(s1)
	length2 := len(s2)

	if length1 > length2 {
		return false
	}

	bytes1 := []rune(s1)
	bytes2 := []rune(s2)

	chars1 := make([]int, 26)
	chars2 := make([]int, 26)

	for i := 0; i < length1; i++ {
		chars1[int(bytes1[i])-97]++
		chars2[int(bytes2[i])-97]++
	}

	for i := 0; i < length2-length1; i++ {
		if match(chars1, chars2) {
			return true
		}
		chars2[int(bytes2[i])-97]--
		chars2[int(bytes2[i+length1])-97]++
	}

	return match(chars1, chars2)
}

func match(c1, c2 []int) bool {
	for i := 0; i < len(c1); i++ {
		if c1[i] != c2[i] {
			return false
		}
	}
	return true
}
