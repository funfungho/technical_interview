package main

/*

https://leetcode.com/problems/range-sum-of-bst/

Easy, BST

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

func rangeSumBST(root *TreeNode, low int, high int) int {
	return dfs(root, low, high)
}

func dfs(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}

	if root.Val < low {
		return dfs(root.Right, low, high)
	} else if root.Val > high {
		return dfs(root.Left, low, high)
	} else if root.Val == low {
		return root.Val + dfs(root.Right, low, high)
	} else if root.Val == high {
		return root.Val + dfs(root.Left, low, high)
	} else {
		return root.Val + dfs(root.Left, low, high) + dfs(root.Right, low, high)
	}
}
