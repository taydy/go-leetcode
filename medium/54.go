package medium

// 	给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
//
//	示例 1:
//		输入:
//			[
//			 [ 1, 2, 3 ],
//			 [ 4, 5, 6 ],
//			 [ 7, 8, 9 ]
//			]
//		输出: [1,2,3,6,9,8,7,4,5]
//
//	示例 2:
//		输入:
//			[
//			  [1, 2, 3, 4],
//			  [5, 6, 7, 8],
//			  [9,10,11,12]
//			]
//		输出: [1,2,3,4,8,12,11,10,9,5,6,7]
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/spiral-matrix
//
func SpiralOrder(matrix [][]int) []int {
	result := make([]int, 0)
	if len(matrix) == 0 {
		return result
	}

	nr := len(matrix) - 1
	nc := len(matrix[0]) - 1
	r1, c1 := 0, 0
	r, c := 0, 0
	for r1 <= nr && c1 <= nc {
		for c = c1; c <= nc; c++ {
			result = append(result, matrix[r1][c])
		}
		for r = r1 + 1; r <= nr; r++ {
			result = append(result, matrix[r][nc])
		}
		if r1 < nr && c1 < nc {
			for c = nc - 1; c > c1; c-- {
				result = append(result, matrix[nr][c])
			}
			for r = nr; r > r1; r-- {
				result = append(result, matrix[r][c1])
			}
		}

		nr--
		nc--
		c1++
		r1++

	}

	return result
}
