package main

/*

https://leetcode.com/problems/sliding-window-maximum/description/

Hard, Monotonic Queue

Time Complexity: O(n)
Space Complexity: O(k)

*/

func maxSlidingWindow(nums []int, k int) []int {
	// brute force: O(kn)
	// monotonic queue: O(n)

	var monoQueue []int // * store index
	results := make([]int, len(nums)-k+1)
	for i, num := range nums {
		// dequeue the front element that's no longer in the window;
		// for loop is not needed, each iteration dequeues one;
		if len(monoQueue) > 0 && monoQueue[0] < i-k+1 {
			monoQueue = monoQueue[1:]
		}

		// maintain a monotonically decreasing queue
		// Within a window, if idxB > idxA and nums[idxB] > nums[idxA],
		// dequeue the tail element
		for len(monoQueue) > 0 && num > nums[monoQueue[len(monoQueue)-1]] {
			// dequque the tail element
			monoQueue = monoQueue[:len(monoQueue)-1]
		}

		// ! enqueue the element that preserves the monotonicity for future use
		monoQueue = append(monoQueue, i)

		if i-k+1 >= 0 {
			results[i-k+1] = nums[monoQueue[0]]
		}
	}

	return results
}
