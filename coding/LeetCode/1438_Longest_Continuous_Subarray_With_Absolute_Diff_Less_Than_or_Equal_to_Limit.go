package main

/*

https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/

Medium, Monotonic Queue (Deque)

Time Complexity: O(n)
Space Complexity: O(n)

Each for loop iteration is amortized O(n) and the deques can grow to size n, where n is the size of nums.
*/

func longestSubarray(nums []int, limit int) int {
	var left, max int

	// ! enqueue to the tail, dequeue from the left
	var increasingMonoQueue, decreasingMonoQueue []int // store value

	for right, num := range nums {
		for len(increasingMonoQueue) > 0 && num < increasingMonoQueue[len(increasingMonoQueue)-1] {
			// dequeue the tail element to maintain monotonically increasing
			increasingMonoQueue = increasingMonoQueue[:len(increasingMonoQueue)-1]
		}
		increasingMonoQueue = append(increasingMonoQueue, num)

		for len(decreasingMonoQueue) > 0 && num > decreasingMonoQueue[len(decreasingMonoQueue)-1] {
			// dequeue the tail element to maintain monotonically decreasing
			decreasingMonoQueue = decreasingMonoQueue[:len(decreasingMonoQueue)-1]
		}
		decreasingMonoQueue = append(decreasingMonoQueue, num)

		// * Trying to keep the constraint: move the left boundary as there's no way the old window can satisfy the constraint.
		// * The left boundary may not align with either of queue's head at first. But as the left boundary keeps moving, it will eventually align with one of the queue's head.
		for decreasingMonoQueue[0]-increasingMonoQueue[0] > limit {
			if nums[left] == increasingMonoQueue[0] {
				increasingMonoQueue = increasingMonoQueue[1:]
			}
			if nums[left] == decreasingMonoQueue[0] {
				decreasingMonoQueue = decreasingMonoQueue[1:]
			}
			// with left moving, the two queues will be dequeued from left, trying to maintain the constraint
			// [2, 1, 4, 7], 5
			// ! this pair of left and right will certainly not satisfy the constraint, so move the left to shrink the window and to see if the new window can satisfy the constraint
			left += 1
		}

		if right-left+1 > max {
			max = right - left + 1
		}
	}

	return max
}
