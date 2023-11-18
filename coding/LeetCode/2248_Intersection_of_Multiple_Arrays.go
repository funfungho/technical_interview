package main

/*

https://leetcode.com/problems/intersection-of-multiple-arrays/

Easy, Hash Table

*/

func intersection(nums [][]int) []int {
	var present [1001]bool
	for _, n := range nums[0] {
		present[n] = true
	}
	for i := 1; i < len(nums); i++ {
		for _, n := range nums[i] {
			if !present[n] {
				present[n] = false
			}
		}
	}
	var ans []int
	for i, n := range present {
		if n {
			ans = append(ans, i)
		}
	}
	return ans
}
