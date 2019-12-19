package hard

import "taydy/go-leetcode/easy"

// 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
//
//	示例:
//
//	输入:
//		[
//  		1->4->5,
//  		1->3->4,
//  		2->6
//		]
//	输出: 1->1->2->3->4->4->5->6

func MergeKLists(lists []*easy.ListNode) *easy.ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	for interval := 1; interval < length; interval *= 2 {
		for i := 0; i < length-interval; i += interval * 2 {
			lists[i] = easy.MergeTwoLists(lists[i], lists[i+interval])
		}
	}
	return lists[0]
}
