package medium

//	判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
//
//	数字 1-9 在每一行只能出现一次。
//	数字 1-9 在每一列只能出现一次。
//	数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/valid-sudoku
//
func IsValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]struct{}, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]struct{})
	}

	columns := make([]map[byte]struct{}, 9)
	for i := 0; i < 9; i++ {
		columns[i] = make(map[byte]struct{})
	}

	boxes := make([]map[byte]struct{}, 9)
	for i := 0; i < 9; i++ {
		boxes[i] = make(map[byte]struct{})
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			t := board[r][c]
			if t == '.' {
				continue
			}
			if _, ok := rows[r][t]; ok {
				return false
			} else {
				rows[r][t] = struct{}{}
			}

			if _, ok := columns[c][t]; ok {
				return false
			} else {
				columns[c][t] = struct{}{}
			}

			box := r/3*3 + c/3
			if _, ok := boxes[box][t]; ok {
				return false
			} else {
				boxes[box][t] = struct{}{}
			}
		}
	}
	return true
}
