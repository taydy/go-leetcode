package medium

//	实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//	如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//	必须原地修改，只允许使用额外常数空间。
//
//	以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
//		1,2,3 → 1,3,2
//		3,2,1 → 1,2,3
//		1,1,5 → 1,5,1
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/next-permutation
//
func NextPermutation(nums []int) {
	i := len(nums) - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	if i >= 0 {
		j := len(nums) - 1
		for ; j >= 0 && nums[i] >= nums[j]; j-- {
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1)
}

// 翻转数组后半部分
func reverse(nums []int, start int) {
	for end := len(nums) - 1; start < end; {
		swap(nums, start, end)
		start++
		end--
	}
}

// 交换数组两个索引的数据
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
