package medium

import "sort"

//	给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//	candidates 中的数字可以无限制重复被选取。
//
//	说明：
//		- 所有数字（包括 target）都是正整数。
//		- 解集不能包含重复的组合。
//
//	示例 1:
//		输入: candidates = [2,3,6,7], target = 7,
//		所求解集为:
//			[
//			  [7],
//			  [2,2,3]
//			]
//	示例 2:
//		输入: candidates = [2,3,5], target = 8,
//		所求解集为:
//			[
//			  [2,2,2,2],
//			  [2,3,3],
//			  [3,5]
//			]
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/combination-sum
//
func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	combinations := make([][]int, 0)
	combinationSumDfs(candidates, nil, &combinations, target, 0)
	return combinations
}

func combinationSumDfs(candidates []int, combination []int, combinations *[][]int, target int, start int) {
	if target == 0 {
		temp := make([]int, len(combination))
		copy(temp, combination)
		*combinations = append(*combinations, temp)
		return
	}

	for i := start; i < len(candidates) && target >= candidates[i]; i++ {
		combinationSumDfs(candidates, append(combination, candidates[i]), combinations, target-candidates[i], i)
	}
}
