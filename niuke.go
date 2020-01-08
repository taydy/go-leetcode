package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	str, err := inputReader.ReadString('\n')

	if err == nil {
		fmt.Println("")
		return
	}

	if len(str) == 0 {
		fmt.Println("")
		return
	}

	bytes := []byte(str)
	start, end := 0, 0
	maxLength := 0
	maxLengthStr := ""
	flagD := false
	for ; end < len(bytes); end++ {
		if bytes[end] == '.' {
			if flagD || (end == 0 || end == len(bytes)-1) || (end > 0 && (bytes[end-1] < 48 || bytes[end-1] > 57)) || (end < len(bytes)-1 && (bytes[end+1] < 48 || bytes[end+1] > 57)) {
				if end-start >= maxLength {
					maxLength = end - start
					maxLengthStr = str[start:end]
				}

				start = end
				flagD = false
				continue
			}
			flagD = true
		} else if bytes[end] == '-' || bytes[end] == '+' {
			if end-start >= maxLength {
				maxLength = end - start
				maxLengthStr = str[start:end]
			}
			start = end
		} else if bytes[end] < 48 || bytes[end] > 57 {
			if end-start >= maxLength {
				maxLength = end - start
				maxLengthStr = str[start:end]
			}

			start = end
			continue
		}
	}
	if end-start > maxLength {
		maxLength = end - start
		maxLengthStr = str[start:end]
	}

	fmt.Println(maxLengthStr)

}
