package main

/*

https://leetcode.com/problems/remove-duplicates-from-sorted-list/

Easy, Linked List

*/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := head
	currNode := head

	for dummy != nil {
		if dummy.Val == currNode.Val {
			if dummy.Next == nil {
				// wrap up, break; remove the trailing same values
				currNode.Next = nil
				break
			} else {
				dummy = dummy.Next
				continue
			}
		} else {
			// break old chain
			currNode.Next = dummy
			// update current node
			currNode = dummy
			dummy = dummy.Next
		}
	}

	return head
}
