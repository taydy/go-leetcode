package medium

//	给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//
//	例如，给定三角形：
//		[
//		     [2],
//		    [3,4],
//		   [6,5,7],
//		  [4,1,8,3]
//		]
//	自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/triangle
//
func MinimumTotal(triangle [][]int) int {

	n := len(triangle)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return triangle[0][0]
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}

	return triangle[0][0]
}
