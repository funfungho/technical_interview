package main

/*

https://leetcode.com/problems/number-of-ways-to-split-array/

Medium, Prefix Sum

*/

func waysToSplitArray(nums []int) int {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]

	// calc prefix sum
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	var ans int
	for i := 0; i < len(nums)-1; i++ {
		if prefixSum[i] >= prefixSum[len(nums)-1]-prefixSum[i] {
			ans++
		}
	}

	return ans
}
