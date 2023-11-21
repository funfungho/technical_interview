package main

/*

https://leetcode.com/problems/ransom-note/

Easy, Hash Table

*/

func canConstruct(ransomNote string, magazine string) bool {
	var letters [26]int
	for _, char := range magazine {
		letters[char-'a']++
	}
	for _, char := range ransomNote {
		if letters[char-'a'] <= 0 {
			return false
		} else {
			letters[char-'a']--
		}
	}
	return true
}
