package medium

import (
	"container/list"
)

// https://leetcode-cn.com/problems/binary-tree-right-side-view/
// 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
// 示例:
//
//		输入: [1,2,3,null,5,null,4]
//		输出: [1, 3, 4]
// 解释:
//
//   		1            <---
// 		  /   \
//		 2     3         <---
// 		 \     \
//  	 5     4       	 <---
//
// 执行结果:
// 执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗 : 2.4 MB , 在所有 Go 提交中击败了 20.00% 的用户
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func rightSideView(root *TreeNode) []int {
	view := make([]int, 0)

	return helper(root, 0, view)
}

func helper(root *TreeNode, level int, view []int) []int {
	if root == nil {
		return view
	}
	if len(view) == level {
		view = append(view, root.Val)
	}
	view = helper(root.Right, level+1, view)
	view = helper(root.Left, level+1, view)
	return view
}

// 非递归
func rightSideView2(root *TreeNode) []int {
	view := make([]int, 0)

	if root == nil {
		return view
	}

	queue := list.New()
	queue.PushFront(root)
	for ; queue.Len() != 0; {
		size := queue.Len()
		for ; size > 0;  {
			node := queue.Front()
			queue.Remove(node)
			if left := node.Value.(*TreeNode).Left; left != nil {
				queue.PushBack(left)
			}
			if right := node.Value.(*TreeNode).Right; right != nil {
				queue.PushBack(right)
			}

			size = size - 1
			if size == 0 {
				view = append(view, node.Value.(*TreeNode).Val)
			}
		}
	}

	return view
}