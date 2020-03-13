package easy

import "sort"

// 	给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//	你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//	示例 1:
//		输入: [3,2,3]
//		输出: 3
//		示例 2:
//	输入: [2,2,1,1,1,2,2]
//		输出: 2
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/majority-element
//
func MajorityElement(nums []int) int {
	return methodII(nums)
}

// 先排序，众数个数因为大于 n/2，所以排序后 n/2 位置一定是众数。
func methodI(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 摩尔投票算法
func methodII(nums []int) int {
	count, candidate := 0, 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}

		if candidate == v {
			count++
		} else {
			count--
		}
	}
	return candidate
}
