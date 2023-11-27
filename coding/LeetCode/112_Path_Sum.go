package main

/*

https://leetcode.com/problems/path-sum/description/

Easy, Tree, DFS

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return true
		}
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// ! wrong implementation

func hasPathSumWrong(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return true
		}
	} else if root.Left != nil {
		// * return result after arriving at the first leaf node
		if hasPathSum(root.Left, targetSum-root.Val) {
			return true
		}
	} else {
		if hasPathSum(root.Right, targetSum-root.Val) {
			return true
		}
	}

	return false
}
