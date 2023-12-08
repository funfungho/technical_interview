package main

/*

https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/

Easy, BST, Inorder Traversal

*/

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	// inorder DFS traversal (left child -> root -> right child)
	var nodeVals []int
	nodeVals = dfs(root, nodeVals)

	min := 100001
	// len(nodeVals) >= 2
	for i := 1; i < len(nodeVals); i++ {
		if nodeVals[i]-nodeVals[i-1] < min {
			min = nodeVals[i] - nodeVals[i-1]
		}
	}
	return min
}

func dfs(root *TreeNode, nodeVals []int) []int {
	if root == nil {
		return nodeVals
	}
	if root.Left != nil {
		nodeVals = dfs(root.Left, nodeVals)
	}
	nodeVals = append(nodeVals, root.Val)
	if root.Right != nil {
		nodeVals = dfs(root.Right, nodeVals)
	}
	return nodeVals
}
