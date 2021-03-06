package easy

//	给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。
//
//	形式上，如果可以找出索引 i+1 < j 且满足 (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ...
//		+ A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]) 就可以将数组三等分。
//
//	示例 1：
//		输出：[0,2,1,-6,6,-7,9,1,2,0,1]
//		输出：true
//		解释：0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
//	示例 2：
//		输入：[0,2,1,-6,6,7,9,-1,2,0,1]
//		输出：false
//	示例 3：
//		输入：[3,3,6,5,-2,2,5,1,-9,4]
//		输出：true
//		解释：3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
//
//	提示：
//		3 <= A.length <= 50000
//		-10^4 <= A[i] <= 10^4
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum
//
func CanThreePartsEqualSum(arr []int) bool {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}

	target := sum / 3

	n, cur, l := len(arr), 0, 0
	for i := 0; i <= 2; i++ {
		sum = 0
		for cur < n {
			sum += arr[cur]
			cur++
			if sum == target {
				l++
				break
			}
		}
	}

	if sum != target || l != 3 {
		return false
	}

	sum = 0
	for cur < n {
		sum += arr[cur]
		cur++
	}

	return sum == 0
}
