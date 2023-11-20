package main

/*

https://leetcode.com/problems/maximum-number-of-balloons/

Easy, Hash Table

*/

func maxNumberOfBalloons(text string) int {
	var records [26]int

	for _, char := range text {
		records[char-'a']++
	}
	min := 10001
	for _, char := range "balon" {
		if min > records[char-'a'] {
			min = records[char-'a']
		}
	}
	if min > records['l'-'a']/2 {
		min = records['l'-'a'] / 2
	}
	if min > records['o'-'a']/2 {
		min = records['o'-'a'] / 2
	}
	return min
}
