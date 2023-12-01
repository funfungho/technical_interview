package main

/*

https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/

Medium, Tree, DFS

*/

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	return diff(root, root.Val, root.Val)
}

func diff(root *TreeNode, max, min int) int {
	if root == nil {
		return max - min
	}

	newMax, newMin := max, min

	if root.Val > max {
		newMax = root.Val
	}

	if root.Val < min {
		newMin = root.Val
	}

	leftMax := diff(root.Left, newMax, newMin)
	rightMax := diff(root.Right, newMax, newMin)
	if leftMax > rightMax {
		return leftMax
	}

	return rightMax
}
