package main

/*

https://leetcode.com/problems/same-tree/description/

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// nil == nil
	if p == q {
		return true
	}

	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		if p.Left == nil && p.Right == nil && q.Left == nil && q.Right == nil && p.Val == q.Val {
			return true
		}
	} else {
		// one nil, the other is non-nil
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
