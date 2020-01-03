package hard

func LongestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}

	max := 0
	for _, v := range nums {
		if _, ok := m[v-1]; ok {
			continue
		}

		tMax := 1
		cur := v
		for {
			if _, ok := m[cur+1]; ok {
				cur = cur + 1
				tMax++
			} else {
				break
			}
		}
		if tMax > max {
			max = tMax
		}

	}

	return max
}
