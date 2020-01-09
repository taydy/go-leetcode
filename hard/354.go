package hard

import (
	"sort"
	"taydy/go-leetcode/medium"
)

//	给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。
//	当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
//
//	请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
//
//	说明:
//		不允许旋转信封。
//
//	示例:
//		输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
//		输出: 3
//		解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/russian-doll-envelopes
//
func MaxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		} else {
			return envelopes[i][1] > envelopes[j][1]
		}
	})

	heights := make([]int, n)
	for i, _ := range envelopes {
		heights[i] = envelopes[i][1]
	}

	return medium.LengthOfLIS(heights)
}
