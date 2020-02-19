package medium

//	给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
//
//	例如，给出 n = 3，生成结果为：
//		[
//		  "((()))",
//		  "(()())",
//		  "(())()",
//			  "()(())",
//		  "()()()"
//		]
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/generate-parentheses
//	解题思路: https://leetcode-cn.com/problems/generate-parentheses/solution/gua-hao-sheng-cheng-by-leetcode/
func GenerateParenthesis(n int) []string {
	results := make([]string, 0)
	generateParenthesisBacktrack(&results, "", 0, 0, n)
	return results
}

func generateParenthesisBacktrack(results *[]string, cur string, open, close, length int) {
	if len(cur) == length<<1 {
		*results = append(*results, cur)
		return
	}

	if open < length {
		generateParenthesisBacktrack(results, cur+"(", open+1, close, length)
	}
	if close < open {
		generateParenthesisBacktrack(results, cur+")", open, close+1, length)
	}
}
