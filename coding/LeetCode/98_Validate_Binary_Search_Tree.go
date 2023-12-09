package main

/*

https://leetcode.com/problems/validate-binary-search-tree/

Medium, BST

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

func isValidBST(root *TreeNode) bool {
	var prevRoot *TreeNode

	// * declare first inorder to call recursively
	var inorderDFS func(root *TreeNode) bool
	inorderDFS = func(root *TreeNode) bool {
		// no need to check nil before invoking recursively
		if root == nil {
			return true
		}

		if !inorderDFS(root.Left) {
			return false
		}
		// by nature, nodes that inorder dfs visit are in monotonically increasing order.
		// think of the visiting process as an array.
		if prevRoot != nil && root.Val <= prevRoot.Val {
			return false
		}
		// prevRoot = root // cannot pass the updated value back to previous call stack
		prevRoot = root
		return inorderDFS(root.Right)
	}

	return inorderDFS(root)
}

/*

min and max are interlacedly passed down to the next call stack

*/
// func isValidBST2(root *TreeNode) bool {
