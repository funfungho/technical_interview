package main

/*

https://leetcode.com/problems/next-greater-element-i/

Easy, Stack, Monotonic Stack

*/

// firstElementLargerOnTheRight
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var monoIncreasingStack []int // store value
	ansMap := make(map[int]int)
	ans := make([]int, len(nums1))

	for i := len(nums2) - 1; i >= 0; i-- {
		// nums2[i]: smaller index, smaller value
		// ! monoIncreasingStack: larger index, larger value: from top to bottom, index is increasing, value is increasing
		/*
				1
			   ---
				2
			   ---
				3
			   ---
				4
			   ---
		*/
		for len(monoIncreasingStack) > 0 && nums2[i] > monoIncreasingStack[len(monoIncreasingStack)-1] {
			// pop
			monoIncreasingStack = monoIncreasingStack[:len(monoIncreasingStack)-1]
		}
		if len(monoIncreasingStack) == 0 {
			// fmt.Println(-1)
			ansMap[nums2[i]] = -1
		} else {
			// * stack top is the larger value with smallest index
			// ! nums[i] is smaller than stack top, so push nums[i] to stack too
			// fmt.Println(monoIncreasingStack[len(monoIncreasingStack)-1])
			ansMap[nums2[i]] = monoIncreasingStack[len(monoIncreasingStack)-1]
		}
		monoIncreasingStack = append(monoIncreasingStack, nums2[i])
	}

	for i, n := range nums1 {
		ans[i] = ansMap[n]
	}

	return ans
}
