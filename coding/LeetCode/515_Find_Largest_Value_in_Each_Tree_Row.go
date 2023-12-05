package main

import "math"

/*

https://leetcode.com/problems/find-largest-value-in-each-tree-row/

Medium, Tree, BFS

*/

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	currLevel := []*TreeNode{root}
	var ans []int
	var nextLevel []*TreeNode
	maxCurrLevel := math.MinInt
	for len(currLevel) > 0 {
		head := currLevel[0]
		if head.Left != nil {
			nextLevel = append(nextLevel, head.Left)
		}
		if head.Right != nil {
			nextLevel = append(nextLevel, head.Right)
		}
		if maxCurrLevel < head.Val {
			maxCurrLevel = head.Val
		}

		// dequeue
		currLevel = currLevel[1:]

		if len(currLevel) == 0 {
			currLevel = nextLevel
			nextLevel = nil
			ans = append(ans, maxCurrLevel)
			maxCurrLevel = math.MinInt
		}
	}

	return ans
}
