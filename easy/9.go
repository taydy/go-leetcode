package easy

// 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
// 示例 1:
//		输入: 121
//		输出: true
//
// 示例 2:
//		输入: -121
//		输出: false
//		解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//
// 示例 3:
//		输入: 10
//		输出: false
//		解释: 从右向左读, 为 01 。因此它不是一个回文数。
//
// 思路:
// 		1. 负数和个数数字为零的整数（0 除外）明显不是回文数;
//		2. 将整数的后半段翻转后如果和前半段相等(长度为偶数的情况下)，或者后半段翻转后去掉个位数字和前半段相等(长度为奇数的情况下)，则表明是回文数
//
// 执行用时 : 32 ms, 在Palindrome Number的Go提交中击败了97.94% 的用户
// 内存消耗 : 5.1 MB, 在Palindrome Number的Go提交中击败了66.21% 的用户
func IsPalindrome(x int) bool {
	// 0 是回文数
	if x == 0 {
		return true
	}
	// 负数和个数数字为零的整数（0 除外）明显不是回文数
	if x < 0 || x%10 == 0 {
		return false
	}

	revertedNumber := 0
	// 依次翻转后半段
	for revertedNumber < x {
		revertedNumber = revertedNumber*10 + x%10
		x = x / 10
	}
	// 后半段翻转后与前半段相等或者后半段翻转后的数字去掉个位和前半段相等，说明是回文数
	return x == revertedNumber || x == revertedNumber/10
}
