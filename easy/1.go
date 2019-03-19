package easy

// 	https://leetcode-cn.com/problems/two-sum/
//	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
//
//	你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//
//	示例:
//
//		给定 nums = [2, 7, 11, 15], target = 9
//
//		因为 nums[0] + nums[1] = 2 + 7 = 9
//		所以返回 [0, 1]
//
//

//
//	在进行迭代并将元素插入到表中的同时，检查表中是否已经存在当前元素所对应的目标元素。如果它存在，则表示已经找到了对应解，并立即将其返回。
//
//	复杂度分析：
//
//		时间复杂度：O(n)， 我们只遍历了包含有 n 个元素的列表一次。在表中进行的每次查找只花费 O(1) 的时间。
//
//		空间复杂度：O(n)， 所需的额外空间取决于哈希表中存储的元素数量，该表最多需要存储 n 个元素。
//

func twoSum(nums []int, target int) []int {
	// 此处字段长度设为数组长度，以空间换取时间，避免大数据时字段频繁扩容影响效率
	dic := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if index, ok := dic[complement]; ok {
			return []int{index, i}
		}
		dic[nums[i]] = i
	}
	return []int{}
}
