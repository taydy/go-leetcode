package medium

// 给你一个 m * n 的矩阵，矩阵中的元素不是 0 就是 1，请你统计并返回其中完全由 1 组成的 正方形 子矩阵的个数。
//
//	示例 1：
//		输入：matrix =
//			[
//			  [0,1,1,1],
//			  [1,1,1,1],
//			  [0,1,1,1]
//			]
//		输出：15
//		解释：
//			边长为 1 的正方形有 10 个。
//			边长为 2 的正方形有 4 个。
//			边长为 3 的正方形有 1 个。
//			正方形的总数 = 10 + 4 + 1 = 15.
//
//	示例 2：
//		输入：matrix =
//			[
//			  [1,0,1],
//			  [1,1,0],
//			  [1,1,0]
//			]
//		输出：7
//		解释：
//			边长为 1 的正方形有 6 个。
//			边长为 2 的正方形有 1 个。
//			正方形的总数 = 6 + 1 = 7.
//
//	提示：
//		1 <= arr.length <= 300
//		1 <= arr[0].length <= 300
//		0 <= arr[i][j] <= 1
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones
//
func CountSquares(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	nr := len(matrix)
	nc := len(matrix[0])

	num := 0
	for r := 0; r < nr; r++ {
		if matrix[r][0] == 1 {
			num++
		}
	}
	for c := 1; c < nc; c++ {
		if matrix[0][c] == 1 {
			num++
		}
	}

	for r := 1; r < nr; r++ {
		for c := 1; c < nc; c++ {
			if matrix[r][c] == 1 {
				matrix[r][c] = intMin(matrix[r-1][c], matrix[r-1][c-1], matrix[r][c-1]) + 1
				num += matrix[r][c]
			}

		}
	}

	return num
}

func intMin(a, b, c int) int {
	min := a
	if min > b {
		min = b
	}
	if min > c {
		min = c
	}
	return min
}
