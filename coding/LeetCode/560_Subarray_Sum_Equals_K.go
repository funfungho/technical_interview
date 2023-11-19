package main

/*

https://leetcode.com/problems/subarray-sum-equals-k/

Medium, Prefix Sum, Hash Table

*/

func subarraySum(nums []int, k int) int {
	ht := make(map[int]int)
	ht[0] = 1
	var ans, prefixSum int

	// ! Focus on the position!
	for _, num := range nums {
		prefixSum += num
		if ht[prefixSum-k] > 0 {
			// There's <ht[prefixSum - k]> number of prefix sums with smaller indices
			// that prefix sum at current index can subtract to get the qualified
			// subarray.
			ans += ht[prefixSum-k]
		}
		// num might be negative, so there might be several indices
		// where a certain prefix sum occurs. Although prefix sum is
		// the same, the positon is not.
		// Mark its existence.
		ht[prefixSum]++
	}

	return ans
}
