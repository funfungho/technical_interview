package main

/*

https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/

Medium, Linked List, Reversal, Fast and Slow Pointers

*/

/*
*
  - Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	fast, slow := head, head
	var halfNodeCount int
	// var prevNode *ListNode
	// * find middle point
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		// prevNode = slow
		slow = slow.Next
		halfNodeCount++
	}
	// node before middle node
	// middleNodeGuide := prevNode // connect to the original tail
	prevMiddleNode := slow // Next is nil

	// ! reverse the right half of the list
	currNode := slow
	nextNode := currNode.Next
	for nextNode != nil {
		next := nextNode.Next
		nextNode.Next = currNode
		currNode = nextNode
		nextNode = next
	}

	// middleNodeGuide.Next = currNode
	prevMiddleNode.Next = nil

	dummyMiddle := currNode
	dummy := head
	var sum int
	// * fast and slow pointers
	for dummyMiddle != nil {
		if sum < dummyMiddle.Val+dummy.Val {
			sum = dummyMiddle.Val + dummy.Val
		}
		dummy = dummy.Next
		dummyMiddle = dummyMiddle.Next
	}

	return sum
}
