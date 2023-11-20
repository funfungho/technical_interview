package main

/*

https://leetcode.com/problems/largest-unique-number/

Easy, Hash Table

*/

func largestUniqueNumber(nums []int) int {
	var ht [1001]int

	for _, num := range nums {
		ht[num]++
	}

	for i := len(ht) - 1; i >= 0; i-- {
		if ht[i] == 1 {
			return i
		}
	}

	return -1
}
