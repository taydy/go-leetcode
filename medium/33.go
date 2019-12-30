package medium

//	假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//	( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//	搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//	你可以假设数组中不存在重复的元素。
//	你的算法时间复杂度必须是 O(log n) 级别。
//
//	示例 1:
//		输入: nums = [4,5,6,7,0,1,2], target = 0
//		输出: 4
//	示例 2:
//		输入: nums = [4,5,6,7,0,1,2], target = 3
//		输出: -1
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
//
func Search1(nums []int, target int) int {
	length := len(nums)

	if length == 0 {
		return -1
	}
	if length == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	// 先查找旋转点
	rotateIndex := findRotateIndex(nums, 0, length-1)

	if nums[rotateIndex] == target {
		return rotateIndex
	} else if nums[length-1] >= target {
		return binarySearch(nums, rotateIndex, len(nums)-1, target)
	} else {
		return binarySearch(nums, 0, rotateIndex, target)
	}

}

func binarySearch(nums []int, left, right int, target int) int {
	for left <= right {
		pivot := left + (right-left)>>1
		if nums[pivot] == target {
			return pivot
		} else if nums[pivot] > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return -1
}

func findRotateIndex(nums []int, left, right int) int {
	if nums[left] < nums[right] {
		return 0
	}

	for left <= right {
		pivot := left + (right-left)>>1
		if nums[pivot] > nums[pivot+1] {
			return pivot + 1
		}
		if nums[pivot] < nums[left] {
			right = pivot - 1
		} else {
			left = pivot + 1
		}

	}
	return 0
}
