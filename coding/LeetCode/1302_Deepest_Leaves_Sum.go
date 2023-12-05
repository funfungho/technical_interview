package main

/*

https://leetcode.com/problems/deepest-leaves-sum/

Medium, Tree

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

func deepestLeavesSum(root *TreeNode) int {
	currLevel := []*TreeNode{root}
	currLevelCopy := []*TreeNode{root}
	var nextLevel []*TreeNode

	var ans int
	for len(currLevel) > 0 {
		head := currLevel[0]
		if head.Left != nil {
			nextLevel = append(nextLevel, head.Left)
		}
		if head.Right != nil {
			nextLevel = append(nextLevel, head.Right)
		}

		// dequeue
		currLevel = currLevel[1:]
		if len(currLevel) == 0 {
			// last level
			if len(nextLevel) == 0 {
				for _, node := range currLevelCopy {
					ans += node.Val
				}
				break
			}

			currLevel = nextLevel
			currLevelCopy = nextLevel
			nextLevel = nil
			ans = 0
		}
	}
	return ans
}
