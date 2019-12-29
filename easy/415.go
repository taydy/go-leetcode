package easy

import (
	"strconv"
	"strings"
)

//	给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//
//	注意：
//		1.num1 和num2 的长度都小于 5100.
//		2.num1 和num2 都只包含数字 0-9.
//		3.num1 和num2 都不包含任何前导零。
//		4.你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/add-strings
//
func AddStrings(num1 string, num2 string) string {
	var sum strings.Builder

	carry := 0
	for s1, s2 := len(num1)-1, len(num2)-1; s1 >= 0 || s2 >= 0; {
		l1, l2 := 0, 0
		if s1 >= 0 {
			l1 = int(num1[s1] - '0')
			s1--
		}
		if s2 >= 0 {
			l2 = int(num2[s2] - '0')
			s2--
		}

		num := carry + l1 + l2
		carry = num / 10
		sum.WriteString(strconv.Itoa(num % 10))
	}

	if carry != 0 {
		sum.WriteString(strconv.Itoa(carry))
	}

	return reverses(sum.String())
}

func reverses(str string) string {
	temp := []byte(str)
	left, right := 0, len(temp)-1
	for left < right {
		temp[left], temp[right] = temp[right], temp[left]
		left++
		right--
	}
	return string(temp)
}
