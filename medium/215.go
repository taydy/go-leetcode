package medium

// 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//	示例 1:
//		输入: [3,2,1,5,6,4] 和 k = 2
//		输出: 5
//
//	示例 2:
//		输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//		输出: 4
//		说明:
//
//	可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

func FindKthLargest(nums []int, k int) int {

	i, start, end := -1, 0, len(nums)-1
	k = len(nums) - k
	for {
		i = partition(nums, start, end)
		if i == k {
			break
		} else if i > k {
			end = i - 1
		} else {
			start = i + 1
		}
	}

	return nums[i]
}

func partition(arr []int, start, end int) int {
	// 选取最后一位当对比数字
	pivot := arr[end]

	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if i != j {
				arr[j], arr[i] = arr[i], arr[j]
			}
			i++
		}
	}
	arr[i], arr[end] = pivot, arr[i]

	return i
}
