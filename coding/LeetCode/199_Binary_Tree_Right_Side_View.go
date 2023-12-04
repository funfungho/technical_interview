package main

/*

https://leetcode.com/problems/binary-tree-right-side-view/description/

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

type NodeWithDepth struct {
	Node  *TreeNode
	Depth int
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ans []int
	lastNode := &NodeWithDepth{
		Node:  root,
		Depth: 0,
	}
	nodeQueue := []*NodeWithDepth{&NodeWithDepth{
		Node:  root,
		Depth: 0,
	}}
	for len(nodeQueue) > 0 {
		head := nodeQueue[0]
		if head.Depth-lastNode.Depth == 1 {
			ans = append(ans, lastNode.Node.Val)
		}
		lastNode = head
		// dequeue
		nodeQueue = nodeQueue[1:len(nodeQueue)]
		// enqueue children
		if head.Node.Left != nil {
			nodeQueue = append(nodeQueue, &NodeWithDepth{
				Node:  head.Node.Left,
				Depth: head.Depth + 1,
			})
		}
		if head.Node.Right != nil {
			nodeQueue = append(nodeQueue, &NodeWithDepth{
				Node:  head.Node.Right,
				Depth: head.Depth + 1,
			})
		}
	}
	ans = append(ans, lastNode.Node.Val)

	return ans
}
