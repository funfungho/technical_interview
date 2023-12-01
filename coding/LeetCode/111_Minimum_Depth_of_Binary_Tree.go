package main

/*

https://leetcode.com/problems/minimum-depth-of-binary-tree/description/

Easy, Tree, DFS

*/

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDepth := 1 + minDepth(root.Left)
	rightDepth := 1 + minDepth(root.Right)

	if leftDepth < rightDepth {
		return leftDepth
	}
	return rightDepth
}
