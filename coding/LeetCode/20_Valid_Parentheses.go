package main

/*

https://leetcode.com/problems/valid-parentheses/submissions/1095345951/

Easy, Stack

Time Complexity: O(n)
Space Complexity: O(n)

*/

func isValid(s string) bool {
	openBrackets := []rune{}

	for _, c := range s {
		switch c {
		case '(', '{', '[':
			openBrackets = append(openBrackets, c)
		default:
			if !isEmpty(openBrackets) {
				top := openBrackets[len(openBrackets)-1]
				if (c == ')' && top == '(') || (c == '}' && top == '{') || (c == ']' && top == '[') {
					// pop
					openBrackets = openBrackets[:len(openBrackets)-1]
				} else {
					// mismatched "[)]" should return false
					return false
				}
			} else {
				// mismatched "]" should return false
				return false
			}
		}
	}

	return isEmpty(openBrackets)
}

func isValid2(s string) bool {
	openBrackets := []rune{}

	matching := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, c := range s {
		if _, isOpenBracket := matching[c]; isOpenBracket {
			openBrackets = append(openBrackets, c)
		} else {
			if isEmpty(openBrackets) {
				// ")"
				return false
			} else {
				top := openBrackets[len(openBrackets)-1]
				if matching[top] != c {
					// "[)]"
					return false
				}
				// pop
				openBrackets = openBrackets[:len(openBrackets)-1]
			}
		}
	}

	return isEmpty(openBrackets)
}

func isEmpty(s []rune) bool {
	return len(s) == 0
}
