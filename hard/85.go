package hard

import "github.com/taydy/go-leetcode/util"

//	给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
//
//	示例:
//		输入:
//			[
//			  ["1","0","1","0","0"],
//			  ["1","0","1","1","1"],
//			  ["1","1","1","1","1"],
//			  ["1","0","0","1","0"]
//			]
//		输出: 6
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/maximal-rectangle
//	题解：https://leetcode-cn.com/problems/maximal-rectangle/solution/zui-da-ju-xing-by-leetcode/
func MaximalRectangle(matrix [][]byte) int {
	row := len(matrix)
	if row == 0 {
		return 0
	}
	col := len(matrix[0])

	left, right, height := make([]int, col), make([]int, col), make([]int, col)
	for i := range right {
		right[i] = col
	}

	maxArea := 0
	for i := 0; i < row; i++ {
		curLeft, curRight := 0, col
		// update height
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				height[j] = height[j] + 1
			} else {
				height[j] = 0
			}
		}

		// update left
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				left[j] = util.MaxInt(left[j], curLeft)
			} else {
				left[j] = 0
				curLeft = j + 1
			}
		}

		// update right
		for j := col - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				right[j] = util.MinInt(right[j], curRight)
			} else {
				right[j] = col
				curRight = j
			}
		}

		// update max area
		for j := 0; j < col; j++ {
			maxArea = util.MaxInt(maxArea, (right[j]-left[j])*height[j])
		}
	}

	return maxArea
}
