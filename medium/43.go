package medium

import (
	"strconv"
	"strings"
)

//	给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
//	示例 1:
//		输入: num1 = "2", num2 = "3"
//		输出: "6"
//	示例 2:
//		输入: num1 = "123", num2 = "456"
//		输出: "56088"
//
//	说明：
//		1.num1 和 num2 的长度小于110。
//		2.num1 和 num2 只包含数字 0-9。
//		3.num1 和 num2 均不以零开头，除非是数字 0 本身。
//		4.不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/multiply-strings
//
func StringMultiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" || len(num1) == 0 || len(num2) == 0 {
		return "0"
	}

	temp := make([]int, len(num1)+len(num2))

	for s1 := len(num1) - 1; s1 >= 0; s1-- {
		for s2 := len(num2) - 1; s2 >= 0; s2-- {
			cur := temp[s1+s2+1] + int(num1[s1]-'0')*int(num2[s2]-'0')
			temp[s1+s2+1] = cur % 10
			temp[s1+s2] += cur / 10
		}
	}

	var sum strings.Builder
	for i := 0; i < len(temp); i++ {
		if sum.Len() == 0 && temp[i] == 0 {
			continue
		}
		sum.WriteString(strconv.Itoa(temp[i]))
	}

	return sum.String()
}
