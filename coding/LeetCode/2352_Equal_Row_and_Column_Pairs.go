package main

import "strings"

/*

https://leetcode.com/problems/equal-row-and-column-pairs/

Medium, Hash Table

*/

func equalPairs(grid [][]int) int {
	// 0: row
	// 1: columnn
	ht := make(map[string][]int)

	matrixSize := len(grid)
	for i := 0; i < matrixSize; i++ {
		var sb strings.Builder
		for j := 0; j < matrixSize; j++ {
			sb.WriteRune(rune(grid[i][j]))
		}
		key := sb.String()
		if _, has := ht[key]; !has {
			ht[key] = make([]int, 2)
		}
		ht[key][0]++
	}

	for i := 0; i < matrixSize; i++ {
		var sb strings.Builder
		for j := 0; j < matrixSize; j++ {
			sb.WriteRune(rune(grid[j][i]))
		}
		key := sb.String()
		if _, has := ht[key]; !has {
			ht[key] = make([]int, 2)
		}
		ht[key][1]++
	}

	var ans int
	for _, v := range ht {
		ans += v[0] * v[1]
	}
	return ans
}
