package medium

// TODO has bug
func longestPalindrome(s string) string {
	length := len(s)
	chars := []rune(s)
	max := 0
	index := make([]int, 128, 128)
	start := 0
	for end := 0; end < length; end++ {

		value := index[chars[end]]
		if value > 0 {
			if curLen := end - value + 1; curLen > max {
				max = curLen
				start = value - 1
			}
		}
		index[chars[end]] = end + 1
	}

	return string(chars[start : start+max+1])
}
