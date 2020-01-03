package medium

import "container/list"

//	给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//	例如：
//		给定二叉树 [3,9,20,null,null,15,7],
//		    3
//		   / \
//		  9  20
//		    /  \
//		   15   7
//		返回锯齿形层次遍历如下：
//		[
//		  [3],
//		  [20,9],
//		  [15,7]
//		]
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
//
func ZigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	l := list.New()
	l.PushBack(root)

	curLevelNum, nextLevelNum, level := 1, 0, 1
	temp := make([]int, 0)
	nextList := list.New()
	for l.Len() > 0 {
		element := l.Front()
		l.Remove(element)

		t := element.Value.(*TreeNode)
		temp = append(temp, t.Val)

		if level%2 == 1 {
			if t.Left != nil {
				nextList.PushFront(t.Left)
				nextLevelNum++
			}
			if t.Right != nil {
				nextList.PushFront(t.Right)
				nextLevelNum++
			}
		} else {
			if t.Right != nil {
				nextList.PushFront(t.Right)
				nextLevelNum++
			}
			if t.Left != nil {
				nextList.PushFront(t.Left)
				nextLevelNum++
			}
		}

		curLevelNum--
		if curLevelNum == 0 {
			level++
			curLevelNum, nextLevelNum = nextLevelNum, 0
			result = append(result, temp)
			temp = make([]int, 0)
			l.PushBackList(nextList)
			nextList = list.New()
		}
	}

	return result
}
