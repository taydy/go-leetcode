package medium

import (
	"container/list"
	"strconv"
	"strings"
)

//	给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
//
//	示例:
//		输入: "25525511135"
//		输出: ["255.255.11.135", "255.255.111.35"]
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/restore-ip-addresses
//
func RestoreIpAddresses(s string) []string {
	output := make([]string, 0)
	segments := list.New()

	backtrack(&output, segments, s, -1, 3)

	return output
}

func backtrack(output *[]string, segments *list.List, s string, prePos int, dots int) {
	maxPos := prePos + 4
	if maxPos > len(s)-1 {
		maxPos = len(s) - 1
	}

	for curPos := prePos + 1; curPos < maxPos; curPos++ {
		segment := s[prePos+1 : curPos+1]
		if valid(segment) {
			segments.PushBack(segment)
			if dots == 1 {
				last := s[curPos+1:]
				if valid(last) {
					segments.PushBack(last)
					*output = append(*output, listJoins(segments, "."))
					segments.Remove(segments.Back())
				}
			} else {
				backtrack(output, segments, s, curPos, dots-1)
			}
			segments.Remove(segments.Back())
		}
	}
}

func valid(segment string) bool {
	length := len(segment)
	if length > 3 {
		return false
	}

	if segment[0] == '0' {
		return length == 1
	}
	v, err := strconv.Atoi(segment)
	return err == nil && v <= 255

}

func listJoins(l *list.List, s string) string {
	var r strings.Builder
	head := l.Front()
	for head != nil {
		r.WriteString(head.Value.(string))
		r.WriteString(s)
		head = head.Next()
	}
	if r.Len() == 0 {
		return ""
	}
	return r.String()[:r.Len()-len(s)]
}
