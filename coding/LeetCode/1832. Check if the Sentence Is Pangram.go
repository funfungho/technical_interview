package main

/*

https://leetcode.com/problems/check-if-the-sentence-is-pangram/

Easy, Hash Table

*/

func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	count := 0
	ht := [26]int{}
	for _, char := range sentence {
		if ht[char-'a'] == 0 {
			count++
			if count == 26 {
				return true
			}
			ht[char-'a']++
		}

	}
	return false
}
