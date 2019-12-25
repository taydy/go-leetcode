package medium

import (
	"math"
)

//	给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
//	假设一个二叉搜索树具有如下特征：
//		1.节点的左子树只包含小于当前节点的数。
//		2.节点的右子树只包含大于当前节点的数。
//		3.所有左子树和右子树自身必须也是二叉搜索树。
//
//	示例 1:
//	输入:
//	    2
//	   / \
//	  1   3
//	输出: true
//
//	示例 2:
//	输入:
//	    5
//	   / \
//	  1   4
//	     / \
//	    3   6
//	输出: false
//
//	解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/validate-binary-search-tree
//	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
func IsValidBST(root *TreeNode) bool {
	// 使用 int64 来避免 int32 超出范围
	return isValidBST(root, math.MinInt64, math.MaxInt64)
}

func isValidBST(root *TreeNode, min, max int64) bool {
	if root == nil {
		return true
	}

	val := int64(root.Val)
	if val <= min || val >= max {
		return false
	}

	return isValidBST(root.Left, min, val) && isValidBST(root.Right, val, max)
}
