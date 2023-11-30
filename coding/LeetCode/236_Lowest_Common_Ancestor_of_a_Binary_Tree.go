package main

/*

https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

Tree, DFS

! The key is to analyze the 3 cases:
1. p and q are in the same subtree: return the subtree root (!important)
2. p and q are in different subtrees: return the root
3. p or q is the root: return the root

*/

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}

	leftSubTreeResult := dfs(root.Left, p, q)
	rightSubTreeResult := dfs(root.Right, p, q)

	if leftSubTreeResult != nil && rightSubTreeResult != nil {
		// find p and q in root's left and sub trees respectively
		return root
	} else if leftSubTreeResult != nil {
		return leftSubTreeResult
	} else {
		return rightSubTreeResult
	}
}

func dfs(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	l := dfs(root.Left, p, q)
	r := dfs(root.Right, p, q)

	if l != nil && r != nil {
		// !
		return root
	} else if l != nil {
		return l
	} else if r != nil {
		return r
	} else {
		return nil
	}
}
