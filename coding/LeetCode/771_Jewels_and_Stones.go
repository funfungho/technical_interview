package main

/*

https://leetcode.com/problems/jewels-and-stones/

Easy, Hash Table

*/

func numJewelsInStones(jewels string, stones string) int {
	var ht [52 + 6]bool
	for _, char := range jewels {
		ht[char-'A'] = true
	}

	var ans int
	for _, char := range stones {
		if ht[char-'A'] {
			ans++
		}
	}
	return ans
}
