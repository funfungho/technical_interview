package main

/*

https://leetcode.com/problems/swap-nodes-in-pairs/

Medium, Linked List, Reversing

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev, curr := head, head.Next
	linkerNode := head
	nodeCount := 2
	// update head
	head = head.Next
	for curr != nil {
		if nodeCount%2 == 0 {
			// swap nodes
			next := curr.Next // save the next node
			curr.Next = prev  // point curr to prev
			linkerNode.Next = curr
			prev.Next = next // prev moves forward
			linkerNode = prev
			curr = next // curr move forward
		} else {
			curr = curr.Next
			prev = prev.Next
		}
		nodeCount++
	}

	return head
}
