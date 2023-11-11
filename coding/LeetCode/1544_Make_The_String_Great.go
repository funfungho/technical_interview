package main

import (
	"fmt"
	"unicode"
)

/*

https://leetcode.com/problems/make-the-string-great/

Easy, Stack

Time Complexity: O(n)
Space Complexity: O(n)

*/

func main() {
	inputs := []string{
		"leEeetcode",
		"abBAcC",
		"s",
	}
	for _, s := range inputs {
		fmt.Printf("%s => %s\n", s, makeGood(s))
	}
}

func makeGood(s string) string {
	var stack []rune

	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}

		top := stack[len(stack)-1]
		// * c != top
		if unicode.ToUpper(c) == unicode.ToUpper(top) && c != top {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	return string(stack)
}
