package main

/*

https://leetcode.com/problems/group-anagrams/

Medium, Hash Table

*/

func groupAnagrams(strs []string) [][]string {
	ht := make(map[[52]rune][]string)

	for _, str := range strs {
		var letters [52]rune
		for _, char := range str {
			letters[char-'a'] = char
			// ["ddddddddddg","dgggggggggg"]
			letters[char-'a'+26]++
		}
		ht[letters] = append(ht[letters], str)
	}

	var ans [][]string
	for _, v := range ht {
		ans = append(ans, v)
	}
	return ans
}
