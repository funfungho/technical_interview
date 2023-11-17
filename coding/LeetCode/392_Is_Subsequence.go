package main

/*

https://leetcode.com/problems/is-subsequence/

Easy, Two Pointers

*/

func isSubsequence(s string, t string) bool {
	var sIdx, tIdx int
	for sIdx < len(s) {
		// t is exhausted, yet s is not matched entirely
		if tIdx > len(t)-1 {
			return false
		}
		if s[sIdx] == t[tIdx] {
			sIdx++
			tIdx++
		} else {
			tIdx++
		}
	}
	return true
}
