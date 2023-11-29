package main

/*

https://leetcode.com/problems/count-good-nodes-in-binary-tree/

Medium, Tree, DFS

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

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}

	// * leaf node
	if root.Left == nil && root.Right == nil {
		if root.Val >= max {
			return 1
		} else {
			return 0
		}
	} else {
		newMax := max
		// update new max
		if root.Val >= max {
			newMax = root.Val
			// * + 1 for current good node
			return 1 + dfs(root.Left, newMax) + dfs(root.Right, newMax)
		}
		// current node is not a good node
		return dfs(root.Left, newMax) + dfs(root.Right, newMax)
	}
}
