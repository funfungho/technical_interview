package main

import (
	"fmt"
	"strings"
)

/*

https://leetcode.com/problems/simplify-path/description/

Medium, Stack

*/

func main() {
	inputs := []string{
		"/././home",
		"/a/b/../home",
		"/a/b/..",
		"/a/b/a../",
		"/..hidden",
	}
	for _, s := range inputs {
		// fmt.Printf("%s => %s\n", s, simplifyPath(s))
		fmt.Printf("%s => %s\n", s, simplifyPath2(s))
	}
}

func simplifyPath2(path string) string {
	components := strings.Split(path, "/")
	var result []string
	for _, comp := range components {
		if comp == "" || comp == "." {
			continue
		} else if comp == ".." {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, comp)
		}
	}

	return "/" + strings.Join(result, "/")
}

func simplifyPath(path string) string {
	var sp []rune

	/*

	   /../../home/ => /home
	   /././home => /home

	   multiple '/' to a single '/';

	   single '.' should be ignored;
	   double '.' should be converted;
	   triple or more '.' remain the unchanged;

	*/

	for j, c := range path {
		if isEmpty(sp) {
			// push the root '/', path guarantees to starts with '/'
			sp = append(sp, c)
		} else {
			if c == '.' {
				sp = append(sp, c)

				// ! handle trailing '.' when iterating over path ends
				if j == len(path)-1 {
					if dots, hasRelativePath := checkRelativePath(sp); hasRelativePath {
						// * absolutize
						// 3 or more consecutive dots are treated as name
						if dots == 1 {
							sp = sp[:len(sp)-1]
						} else if dots == 2 {
							if len(sp) == 3 {
								// "/..", only pop last two dots
								sp = sp[:len(sp)-2]
							} else {
								// "/a/b/..", pop "/.." first
								sp = sp[:len(sp)-3]
								// keep popping until '/'
								for i := len(sp) - 1; i >= 0; i-- {
									if sp[i] == '/' {
										// preserve '/'
										sp = sp[:i+1]
										break
									}
								}
							}
						}

					}
				}
			} else if c == '/' {
				// ! handle '.' before push other non-dot characters
				dots, hasRelativePath := checkRelativePath(sp)
				if !hasRelativePath {
					// handle '/' specifically: multiple '/' to a single '/'
					if top(sp) != '/' {
						sp = append(sp, c)
					}
				} else {
					// * absolutize
					if dots == 1 {
						// pop redundant dot
						sp = sp[:len(sp)-1]

						// handle '/' specifically: multiple '/' to a single '/'
						if top(sp) != '/' {
							sp = append(sp, c)
						}
					} else if dots == 2 {
						if len(sp) == 3 {
							// "/..", only pop last two dots
							sp = sp[:len(sp)-2]
						} else {
							// "/a/b/..", pop "/.." first
							sp = sp[:len(sp)-3]
							// keep popping until '/'
							for i := len(sp) - 1; i >= 0; i-- {
								if sp[i] == '/' {
									// preserve '/'
									sp = sp[:i+1]
									break
								}
							}
						}

						// handle '/' specifically: multiple '/' to a single '/'
						if top(sp) != '/' {
							sp = append(sp, c)
						}
					} else {
						// 3 or more consecutive dots are treated as name
						sp = append(sp, c)
					}
				}
			} else {
				// non-dot character and non-slash character clear the overhead of checking relative path
				sp = append(sp, c)
			}
		}
	}

	if len(sp) > 1 && sp[len(sp)-1] == '/' {
		// remove trailing '/'
		sp = sp[:len(sp)-1]
	}

	return string(sp)
}

func isEmpty(s []rune) bool {
	return len(s) == 0
}

func top(s []rune) rune {
	return s[len(s)-1]
}

/*
check trigger condition for relative path:
 1. iterating over path ends;
 2. encounter a non-dot character;
*/
func checkRelativePath(s []rune) (int, bool) {
	dots := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '.' {
			if s[i] != '/' {
				dots = 0
			}
			break
		}
		dots++
	}
	return dots, dots != 0
}
