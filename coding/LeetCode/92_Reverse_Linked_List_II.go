package main

/*

https://leetcode.com/problems/reverse-linked-list-ii/

Medium, Linked List, Reversal

*/

/*
*
  - Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := head
	count := 1
	link2ReversedHead := dummy // link to the new "head"
	for dummy.Next != nil && count < left {
		link2ReversedHead = dummy
		dummy = dummy.Next
		count++
	}

	// start reversal from dummy node
	reversedTail := dummy  // link to the the remaining un-reversed chain
	nextNode := dummy.Next // won't be nil
	currNode := dummy
	nodesToReverse := right - left
	var reversed int
	for nextNode != nil && reversed < nodesToReverse {
		next := nextNode.Next
		nextNode.Next = currNode
		currNode = nextNode
		nextNode = next
		reversed++
	}
	link2ReversedHead.Next = currNode
	reversedTail.Next = nextNode

	// *
	if left == 1 {
		return currNode
	}

	return head
}
