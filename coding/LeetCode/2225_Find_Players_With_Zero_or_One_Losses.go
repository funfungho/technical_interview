package main

/*

https://leetcode.com/problems/find-players-with-zero-or-one-losses/

Medium, Hash Table

*/

func findWinners(matches [][]int) [][]int {
	var records [100001]int
	// -1 (played no game)
	// 0: win all
	// 1: lose 1
	// > 1: lose more than 1
	for i := range records {
		records[i] = -1
	}

	for _, match := range matches {
		if records[match[0]] == -1 {
			// first win
			records[match[0]] = 0
		}
		if records[match[1]] == -1 {
			// first lose
			records[match[1]] = 1
		} else {
			records[match[1]]++
		}
	}
	var winAll, loseOne []int
	for i, record := range records {
		if record == 0 {
			winAll = append(winAll, i)
		} else if record == 1 {
			loseOne = append(loseOne, i)
		}
	}

	return [][]int{winAll[:], loseOne[:]}
}
