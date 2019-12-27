package medium

//	给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。
//	你可以假设二维矩阵的四个边缘都被水包围着。
//
//	找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。
//
//	示例 1:
//		[[0,0,1,0,0,0,0,1,0,0,0,0,0],
//		 [0,0,0,0,0,0,0,1,1,1,0,0,0],
//		 [0,1,1,0,1,0,0,0,0,0,0,0,0],
//		 [0,1,0,0,1,1,0,0,1,0,1,0,0],
//		 [0,1,0,0,1,1,0,0,1,1,1,0,0],
//		 [0,0,0,0,0,0,0,0,0,0,1,0,0],
//		 [0,0,0,0,0,0,0,1,1,1,0,0,0],
//		 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
//	对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。
//
//	示例 2:
//	[[0,0,0,0,0,0,0,0]]
//	对于上面这个给定的矩阵, 返回 0。
//
//	注意: 给定的矩阵grid 的长度和宽度都不超过 50。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/max-area-of-island
//
func MaxAreaOfIsland(grid [][]int) int {
	if grid == nil {
		return 0
	}

	nr := len(grid)

	if nr == 0 {
		return 0
	}

	nc := len(grid[0])
	maxArea := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				area := maxAreaHelper(grid, r, c, nr, nc, 0)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func maxAreaHelper(grid [][]int, r int, c int, nr int, nc int, area int) int {
	if r < 0 || c < 0 || r >= nr || c >= nc || grid[r][c] == 0 {
		return area
	}
	grid[r][c] = 0
	area++
	area = maxAreaHelper(grid, r-1, c, nr, nc, area)
	area = maxAreaHelper(grid, r+1, c, nr, nc, area)
	area = maxAreaHelper(grid, r, c-1, nr, nc, area)
	area = maxAreaHelper(grid, r, c+1, nr, nc, area)
	return area
}
