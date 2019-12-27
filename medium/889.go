package medium

//	返回与给定的前序和后序遍历匹配的任何二叉树。
//
//	 pre 和 post 遍历中的值是不同的正整数。
//
//	示例：
//
//		输入：pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
//		输出：[1,2,3,4,5,6,7]
//
//	提示：
//		1 <= pre.length == post.length <= 30
//		pre[] 和 post[] 都是 1, 2, ..., pre.length 的排列
//		每个输入保证至少有一个答案。如果有多个答案，可以返回其中一个。
//
//	来源：力扣（LeetCode）
//	链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal
//
func constructFromPrePost(pre []int, post []int) *TreeNode {
	pLength := len(pre)
	if pLength == 0 {
		return nil
	}

	root := &TreeNode{
		Val: pre[0],
	}

	if pLength == 1 {
		return root
	}

	mid := -1
	for i, v := range post {
		if v == pre[1] {
			mid = i
			break
		}
	}

	if mid == -1 {
		return root
	}

	root.Left = constructFromPrePost(pre[1:mid+2], post[:mid+1])
	root.Right = constructFromPrePost(pre[mid+2:], post[mid+1:len(post)-1])
	return root
}
