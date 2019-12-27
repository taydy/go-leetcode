package medium

//	在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
//
//	示例:
//
//	输入:
//		1 0 1 0 0
//		1 0 1 1 1
//		1 1 1 1 1
//		1 0 0 1 0
//	输出: 4
//
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/maximal-square
//
func MaximalSquare(matrix [][]byte) int {

	if len(matrix) == 0 {
		return 0
	}
	nr := len(matrix)
	nc := len(matrix[0])
	max := 0

	for r := 0; r < nr; r++ {
		if matrix[r][0] == '1' {
			max = 1
			break
		}
	}

	if max == 0 {
		for c := 0; c < nc; c++ {
			if matrix[0][c] == '1' {
				max = 1
				break
			}
		}
	}

	for r := 1; r < nr; r++ {
		for c := 1; c < nc; c++ {
			if matrix[r][c] == '1' {
				matrix[r][c] = byte(int(mins(matrix[r-1][c], matrix[r-1][c-1], matrix[r][c-1])) + 1)
				if max < int(matrix[r][c])-48 {
					max = int(matrix[r][c]) - 48
				}
			}

		}
	}
	return max * max
}

func mins(a, b, c byte) byte {
	min := a
	if min > b {
		min = b
	}
	if min > c {
		min = c
	}
	return min
}
