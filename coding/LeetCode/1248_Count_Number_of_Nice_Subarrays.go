package main

/*

https://leetcode.com/problems/count-number-of-nice-subarrays/

Medium, Prefix Sum, Hash Table

*/

func numberOfSubarrays(nums []int, k int) int {
	var oddNumberCount int
	oddNumberCountMap := make(map[int]int)
	var ans int

	oddNumberCountMap[0] = 1

	for _, num := range nums {
		oddNumberCount += num % 2
		if v, exists := oddNumberCountMap[oddNumberCount-k]; exists {
			ans += v
		}
		oddNumberCountMap[oddNumberCount]++
	}

	return ans
}
