package main

/*

https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/

Easy, Stack

Time Complexity: O(n)
Space Complexity: O(n)

*/

func removeDuplicates(s string) string {
	var removed []rune

	for _, c := range s {
		if isEmpty(removed) || removed[len(removed)-1] != c {
			removed = append(removed, c)
		} else {
			// pop
			removed = removed[:len(removed)-1]
		}
	}

	return string(removed)
}

func isEmpty(s []rune) bool {
	return len(s) == 0
}
