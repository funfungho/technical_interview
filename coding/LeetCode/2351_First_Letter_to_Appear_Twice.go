package main

/*

https://leetcode.com/problems/first-letter-to-appear-twice/

Easy, Hash Table

*/

func repeatedCharacter(s string) byte {
	ht := make(map[byte]bool)
	for _, char := range s {
		if _, has := ht[byte(char)]; has {
			return byte(char)
		}
		ht[byte(char)] = true
	}
	return 1
}
