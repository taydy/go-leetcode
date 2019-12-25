package hard

// 给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。
//
//	你可以对一个单词进行如下三种操作：
//		1.插入一个字符
//		2.删除一个字符
//		3.替换一个字符
//
// 示例 1:
//		输入: word1 = "horse", word2 = "ros"
//		输出: 3
//		解释:
//			horse -> rorse (将 'h' 替换为 'r')
//			rorse -> rose (删除 'r')
//			rose -> ros (删除 'e')
//
//示例 2:
//		输入: word1 = "intention", word2 = "execution"
//		输出: 5
//		解释:
//			intention -> inention (删除 't')
//			inention -> enention (将 'i' 替换为 'e')
//			enention -> exention (将 'n' 替换为 'x')
//			exention -> exection (将 'n' 替换为 'c')
//			exection -> execution (插入 'u')

func MinDistance(word1 string, word2 string) int {
	c1, c2 := []byte(word1), []byte(word2)

	if len(c1) == 0 {
		return len(c2)
	}
	if len(c2) == 0 {
		return len(c1)
	}

	return minEditBT2(c1, len(c1), c2, len(c2))
	//return minEditBT1(0, 0, c1, c2, len(c1), len(c2), 0, make(map[string]int))
}

// 方法 1
// 如果字符串很长,此方法计算会超时
func minEditBT1(i int, j int, c1 []byte, c2 []byte, len1 int, len2 int, editBT int, m map[string]int) int {
	if i == len1 || j == len2 {
		if i < len1 {
			return editBT + (len1 - i)
		}
		if j < len2 {
			return editBT + (len2 - j)
		}
		return editBT
	}

	// 两字符匹配
	if c1[i] == c2[j] {
		return minEditBT1(i+1, j+1, c1, c2, len1, len2, editBT, m)
	} else {
		temp1 := minEditBT1(i+1, j, c1, c2, len1, len2, editBT+1, m)
		temp2 := minEditBT1(i, j+1, c1, c2, len1, len2, editBT+1, m)
		temp3 := minEditBT1(i+1, j+1, c1, c2, len1, len2, editBT+1, m)

		editBT = temp1
		if editBT > temp2 {
			editBT = temp2
		}
		if editBT > temp3 {
			editBT = temp3
		}
		return editBT
	}
}

// 方法 2
// 先依次将每一步计算好,填充进二维数组中,这样可以避免某一步被多次计算
// 如果：a[i]!=b[j]，那么：min_edist(i, j)就等于：min(min_edist(i-1,j)+1, min_edist(i,j-1)+1, min_edist(i-1,j-1)+1)
// 如果：a[i]==b[j]，那么：min_edist(i, j)就等于：min(min_edist(i-1,j)+1, min_edist(i,j-1)+1，min_edist(i-1,j-1))其中，min表示求三数中的最小值
func minEditBT2(c1 []byte, len1 int, c2 []byte, len2 int) int {
	minDist := make([][]int, len1)
	for i := 0; i < len1; i++ {
		minDist[i] = make([]int, len2)
	}

	// 初始化第0行:a[0..0]与b[0..j]的编辑距离
	for j := 0; j < len2; j++ {
		if c1[0] == c2[j] {
			minDist[0][j] = j
		} else if j != 0 {
			minDist[0][j] = minDist[0][j-1] + 1
		} else {
			minDist[0][j] = 1
		}
	}

	// 初始化第0列:a[0..i]与b[0..0]的编辑距离
	for i := 0; i < len1; i++ {
		if c2[0] == c1[i] {
			minDist[i][0] = i
		} else if i != 0 {
			minDist[i][0] = minDist[i-1][0] + 1
		} else {
			minDist[i][0] = 1
		}
	}

	// 按行填表
	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			if c1[i] == c2[j] {
				minDist[i][j] = min(minDist[i-1][j]+1, minDist[i][j-1]+1, minDist[i-1][j-1])
			} else {
				minDist[i][j] = min(minDist[i-1][j]+1, minDist[i][j-1]+1, minDist[i-1][j-1]+1)
			}
		}
	}

	return minDist[len1-1][len2-1]
}

func min(x, y, z int) int {
	minV := x
	if minV > y {
		minV = y
	}
	if minV > z {
		return z
	}
	return minV
}
