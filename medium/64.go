package medium

//	给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//	说明：每次只能向下或者向右移动一步。
//
//	示例:
//		输入:
//			[
//			  [1,3,1],
//			  [1,5,1],
//			  [4,2,1]
//			]
//		输出: 7
//		解释: 因为路径 1→3→1→1→1 的总和最小。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-path-sum
//
func MinPathSum(grid [][]int) int {
	nr := len(grid)
	nc := len(grid[0])
	for r := 1; r < nr; r++ {
		grid[r][0] = grid[r][0] + grid[r-1][0]
	}

	for c := 1; c < nc; c++ {
		grid[0][c] = grid[0][c] + grid[0][c-1]
	}

	for r := 1; r < nr; r++ {
		for c := 1; c < nc; c++ {
			grid[r][c] = min(grid[r-1][c], grid[r][c-1]) + grid[r][c]
		}
	}

	return grid[nr-1][nc-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
