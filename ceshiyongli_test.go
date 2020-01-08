package main

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	str := "+012345.67890.967890111111a"
	fmt.Printf("%s, %d", str, len(str))
	bytes := []byte(str)
	start, end := 0, 0
	maxLength := 0
	maxLengthStr := ""
	flagD := false
	lastDIndex := 0
	for ; end < len(bytes); end++ {
		if bytes[end] == '.' {
			if flagD || (end == 0 || end == len(bytes)-1) || (end > 0 && (bytes[end-1] < 48 || bytes[end-1] > 57)) || (end < len(bytes)-1 && (bytes[end+1] < 48 || bytes[end+1] > 57)) {
				if end-start >= maxLength {
					maxLength = end - start
					maxLengthStr = str[start:end]
				}

				if lastDIndex != start && end-lastDIndex >= maxLength {
					maxLength = end - lastDIndex
					maxLengthStr = str[lastDIndex:end]
					lastDIndex = end + 1
				}

				start = end + 1
				flagD = false
				continue
			}
			lastDIndex = end + 1 // 记录上一次小数点位置
			flagD = true
		} else if bytes[end] == '-' || bytes[end] == '+' {
			if end-start >= maxLength {
				maxLength = end - start
				maxLengthStr = str[start:end]
			}
			lastDIndex = end
			start = end
		} else if bytes[end] < 48 || bytes[end] > 57 {
			if end-start >= maxLength {
				maxLength = end - start
				maxLengthStr = str[start:end]
			}

			if lastDIndex != start && end-lastDIndex >= maxLength {
				maxLength = end - lastDIndex
				maxLengthStr = str[lastDIndex:end]
				lastDIndex = end + 1
			}

			start = end + 1
			lastDIndex = end + 1
			continue
		}
	}

	if end-start >= maxLength {
		maxLength = end - start
		maxLengthStr = str[start:end]
	}

	if lastDIndex != start && end-lastDIndex >= maxLength {
		maxLength = end - lastDIndex
		maxLengthStr = str[lastDIndex:end]
	}

	if maxLengthStr == "+" || maxLengthStr == "-" {
		maxLengthStr = ""
	}

	t.Log(maxLengthStr)
}
