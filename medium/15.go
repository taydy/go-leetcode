package medium

import "sort"

//	给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//	注意：答案中不可以包含重复的三元组。
//
//	例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//	满足要求的三元组集合为：
//	[
//	  [-1, 0, 1],
//	  [-1, -1, 2]
//	]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/3sum
//
func ThreeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return [][]int{}
	}

	// 先排序
	sort.Ints(nums)

	result := make([][]int, 0)

	for i, v := range nums {
		if v > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := length - 1

		for l < r {
			sum := v + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{v, nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return result
}
