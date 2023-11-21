package main

/*

https://leetcode.com/problems/longest-substring-without-repeating-characters/

Medium, Hash Table, Sliding Window

!important

*/

func lengthOfLongestSubstring(s string) int {
	var max int
	// value: count
	ht := make(map[rune]int)
	var left, right int
	var char rune

	for right, char = range s {
		ht[char]++
		// ! move left past the duplicate one
		// ! use for loop to handle all cases
		for ht[char] > 1 {
			charAtLeft := s[left]
			// move left
			ht[rune(charAtLeft)]--
			left++
		}
		if max < right-left+1 {
			max = right - left + 1
		}
	}

	return max
}

/*
  "abcabcbb"
  "bbbbb"
  "pwwkew"
  " "
  ""
  "au"
  "aau"
  "dvdf"
  "abcdeff"
  "abba"
  "tmmzuxt"
*/
