package main

/*
https://leetcode.com/problems/counting-elements/

Easy, Hash Table
*/

func countElements(arr []int) int {
	ht := make(map[int]struct{})
	for _, n := range arr {
		ht[n] = struct{}{}
	}
	var count int
	for _, n := range arr {
		_, hasN := ht[n]
		_, hasNPlus := ht[n+1]
		if hasN && hasNPlus {
			count++
		}
	}

	return count
}
