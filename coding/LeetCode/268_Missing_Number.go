package main

/*

https://leetcode.com/problems/missing-number/

Easy, Hash Table

*/

func missingNumber(nums []int) int {
	total := len(nums) * (len(nums) + 1) / 2

	for _, num := range nums {
		total -= num
	}

	return total
}
