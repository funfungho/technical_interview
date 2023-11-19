package main

/*

https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/

Easy, Hash Table

*/

func areOccurrencesEqual(s string) bool {
	ht := [26]int{}
	for _, char := range s {
		ht[char-'a']++
	}

	var count int
	for _, f := range ht {
		if f != 0 {
			if count == 0 {
				count = f
			} else if f != count {
				return false
			}
		}
	}
	return true
}
