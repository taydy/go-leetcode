package hard

//	给定一个未排序的整数数组，找出最长连续序列的长度。
//	要求算法的时间复杂度为 O(n)。
//
//	示例:
//		输入: [100, 4, 200, 1, 3, 2]
//		输出: 4
//		解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
//
func LongestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}

	max := 0
	for _, v := range nums {
		if _, ok := m[v-1]; ok {
			continue
		}

		tMax := 1
		cur := v
		for {
			if _, ok := m[cur+1]; ok {
				cur = cur + 1
				tMax++
			} else {
				break
			}
		}
		if tMax > max {
			max = tMax
		}

	}

	return max
}
