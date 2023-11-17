package main

/*

https://leetcode.com/problems/running-sum-of-1d-array/

Easy, Prefix Sum

*/

func runningSum(nums []int) []int {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	return prefixSum
}
