package main

import "math"

/*

https://leetcode.com/problems/closest-binary-search-tree-value/

Easy, Binary Search Tree

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

func closestValue(root *TreeNode, target float64) int {
	gap := math.MaxFloat64
	ans := root.Val
	dfs(root, target, &gap, &ans)
	return ans
}

func dfs(root *TreeNode, target float64, gap *float64, ans *int) {
	if root == nil {
		return
	}

	if float64(root.Val) > target {
		if float64(root.Val)-target < *gap {
			*ans = root.Val
			*gap = float64(root.Val) - target
		} else if float64(root.Val)-target == *gap && root.Val < *ans {
			*ans = root.Val
		}
		dfs(root.Left, target, gap, ans)
	} else if float64(root.Val) < target {
		if target-float64(root.Val) < *gap {
			*ans = root.Val
			*gap = target - float64(root.Val)
		} else if target-float64(root.Val) == *gap && root.Val < *ans {
			*ans = root.Val
		}
		dfs(root.Right, target, gap, ans)
	} else {
		*ans = root.Val
	}
}
