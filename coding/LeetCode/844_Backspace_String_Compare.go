package main

/*

https://leetcode.com/problems/backspace-string-compare/

Easy, Stack

Time Complexity: O(n)
Space Complexity: O(n)

*/

func backspaceCompare(s string, t string) bool {
	return backspaced(s) == backspaced(t)
}

func backspaced(str string) string {
	var s []rune

	for _, c := range str {
		if c != '#' {
			// push
			s = append(s, c)
		} else if !isEmpty(s) {
			// pop
			s = s[:len(s)-1]
		}
	}

	return string(s)
}

func isEmpty(s []rune) bool {
	return len(s) == 0
}
