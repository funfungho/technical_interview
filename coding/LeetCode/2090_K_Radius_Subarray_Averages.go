package main

/*

https://leetcode.com/problems/k-radius-subarray-averages/

Medium, Prefix Sum

*/

func getAverages(nums []int, k int) []int {
	if k == 0 {
		return nums
	}

	prefixSum := make([]int, len(nums)+1)

	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}

	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	for i := k; i <= len(nums)-k-1; i++ {
		ans[i] = (prefixSum[i+k+1] - prefixSum[i-k]) / (2*k + 1)
	}

	return ans
}
