package medium

//	给定一个没有重复数字的序列，返回其所有可能的全排列。
//
//	示例:
//		输入: [1,2,3]
//		输出:
//			[
//			  [1,2,3],
//			  [1,3,2],
//			  [2,1,3],
//			  [2,3,1],
//			  [3,1,2],
//			  [3,2,1]
//			]
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/permutations
//
func Permute(nums []int) [][]int {
	output := make([][]int, 0)

	permuteBacktrack(&output, nums, 0, len(nums))

	return output
}

func permuteBacktrack(output *[][]int, nums []int, first int, length int) {
	if first == length {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*output = append(*output, temp)
	}

	for i := first; i < length; i++ {
		nums[first], nums[i] = nums[i], nums[first]
		permuteBacktrack(output, nums, first+1, length)
		nums[first], nums[i] = nums[i], nums[first]
	}
}
