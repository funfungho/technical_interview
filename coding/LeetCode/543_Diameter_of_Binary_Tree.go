package main

/*

https://leetcode.com/problems/diameter-of-binary-tree/

Easy (not easy), Tree, DFS

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

func diameterOfBinaryTree(root *TreeNode) int {
	var currMax int
	maxDepth(root, &currMax)
	return currMax
}

func maxDepth(root *TreeNode, currMax *int) int {
	if root == nil {
		return 0
	}

	// leaf node
	if root.Left == nil && root.Right == nil {
		return 1
	}

	maxLeftDepth := maxDepth(root.Left, currMax)
	maxRightDepth := maxDepth(root.Right, currMax)
	// https://leetcode.com/problems/diameter-of-binary-tree/Figures/543/543.png
	// case 1 in the above image
	if maxLeftDepth+maxRightDepth > *currMax {
		*currMax = maxLeftDepth + maxRightDepth
	}

	// ! +1 means the edge between current root and its parent, for usage in the recursion of its parent
	// case 1 in the above image
	if maxLeftDepth > maxRightDepth {
		return maxLeftDepth + 1
	}
	return maxRightDepth + 1
}
