package main

/*

https://leetcode.com/problems/daily-temperatures/

Medium, Monotonic Stack

Time Complexity: O(2n) => O(n)
Brute Force: O(n^2)

*/

func dailyTemperatures(temperatures []int) []int {
	/*
		Maintain a monotonically **decreasing** stack. The stack stores the element
		index.

		Every time it encounters an item that is larger than the top element,
		the stack pops the top element in a loop until the top element is smaller
		than the new element. Then the new  element is pushed onto the stack.

		! Every time an item is popped from the stack, the result at that index can be calculated.
	*/

	var indexStack []int
	results := make([]int, len(temperatures))
	for i, t := range temperatures {
		for len(indexStack) > 0 && t > temperatures[indexStack[len(indexStack)-1]] {
			topIndex := indexStack[len(indexStack)-1]
			// calculate the gap
			results[topIndex] = i - topIndex
			// pop
			indexStack = indexStack[:len(indexStack)-1]
		}
		indexStack = append(indexStack, i)
	}

	return results
}
