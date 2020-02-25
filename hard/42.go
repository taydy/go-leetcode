package hard

import "github.com/taydy/go-leetcode/util"

//	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//	上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。
//
//	示例:
//		输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//		输出: 6
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/trapping-rain-water
//

// 每一次循环将所有的数据减 1,如果某个高度小于 0 则说明此处可以接到一格的雨水,
// 去除头部和尾部小于等于 0 的位置.
func Trap(height []int) int {
	max := 0
	for len(height) > 0 {
		for len(height) > 0 && height[0] <= 0 {
			height = height[1:]
		}
		for len(height) > 0 && height[len(height)-1] <= 0 {
			height = height[:len(height)-1]
		}

		for i, v := range height {
			if i != 0 && i != len(height)-1 {
				if v <= 0 {
					max++
				}
			}
			height[i]--
		}
	}
	return max
}

// 用两个数组分别记录位置 i 的左侧最大高度和右侧最大高度, 两者的最小值减去 i 本身的高度就是 i 位置可以接的雨水.
func Trap1(height []int) int {
	n := len(height)

	if n <= 2 {
		return 0
	}

	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0], rightMax[n-1] = height[0], height[n-1]

	for i := 1; i < n; i++ {
		leftMax[i] = util.MaxInt(leftMax[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = util.MaxInt(rightMax[i+1], height[i])
	}

	max := 0
	for i, v := range height {
		max += util.MinInt(leftMax[i], rightMax[i]) - v
	}

	return max
}
