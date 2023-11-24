package main

/*

https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

Easy, Tree, DFS

*/

/*
*
  - Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// post order traversal
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	max := 1
	if leftDepth < rightDepth {
		max += rightDepth
	} else {
		max += leftDepth
	}

	return max
}
