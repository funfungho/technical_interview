package main

/*

https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	currLevel := []*TreeNode{root}
	var currLevelVals []int
	var nextLevel []*TreeNode
	var level int
	var head *TreeNode
	var ans [][]int

	for len(currLevel) > 0 {
		head = currLevel[0]
		currLevel = currLevel[1:]
		currLevelVals = append(currLevelVals, head.Val)

		if head.Left != nil {
			nextLevel = append(nextLevel, head.Left)
		}
		if head.Right != nil {
			nextLevel = append(nextLevel, head.Right)
		}

		if len(currLevel) == 0 {
			if level%2 != 0 {
				for i, j := 0, len(currLevelVals)-1; i < j; i, j = i+1, j-1 {
					currLevelVals[i], currLevelVals[j] = currLevelVals[j], currLevelVals[i]
				}
			}
			ans = append(ans, currLevelVals)
			level++
			if len(nextLevel) == 0 {
				break
			}
			currLevel = nextLevel
			currLevelVals = nil
			nextLevel = nil
		}
	}

	return ans
}
