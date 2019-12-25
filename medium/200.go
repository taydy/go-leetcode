package medium

//	给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，
//	并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
//
//	示例 1:
//		输入:
//			11110
//			11010
//			11000
//			00000
//		输出: 1
//
//	示例 2:
//		输入:
//			11000
//			11000
//			00100
//			00011
//		输出: 3
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/number-of-islands
//
// 思路:
//		线性扫描整个二维网格，如果一个结点包含 1，则以其为根结点启动深度优先搜索。在深度优先搜索过程中，
//		每个访问过的结点被标记为 0。计数启动深度优先搜索的根结点的数量，即为岛屿的数量。
//
func NumIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	nr := len(grid)
	nc := len(grid[0])

	landsNum := 0

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				landsNum++
				dfs(grid, r, c, nr, nc)
			}
		}
	}

	return landsNum
}

func dfs(grid [][]byte, r, c, nr, nc int) {

	if r < 0 || r >= nr || c < 0 || c >= nc || grid[r][c] == '0' {
		return
	}

	grid[r][c] = '0'
	dfs(grid, r-1, c, nr, nc)
	dfs(grid, r+1, c, nr, nc)
	dfs(grid, r, c-1, nr, nc)
	dfs(grid, r, c+1, nr, nc)
}
