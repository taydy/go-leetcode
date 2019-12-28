package medium

//	给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
//
//	示例:
//
//	输入: 3
//	输出:
//		[
//		 [ 1, 2, 3 ],
//		 [ 8, 9, 4 ],
//		 [ 7, 6, 5 ]
//		]
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/spiral-matrix-ii
//
func GenerateMatrix(n int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}

	nr, nc := n-1, n-1
	r1, c1, r, c := 0, 0, 0, 0
	cur := 1
	for r1 <= nr && c1 <= nc {
		for c = c1; c <= nc; c++ {
			matrix[r1][c] = cur
			cur++
		}
		for r = r1 + 1; r <= nr; r++ {
			matrix[r][nc] = cur
			cur++
		}

		for c = nc - 1; c > c1; c-- {
			matrix[nr][c] = cur
			cur++
		}
		for r = nr; r > r1; r-- {
			matrix[r][c1] = cur
			cur++
		}

		nr--
		nc--
		c1++
		r1++
	}

	return matrix
}
