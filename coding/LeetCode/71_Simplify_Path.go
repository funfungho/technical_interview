package main

import "fmt"

func main() {
	inputs := []string{
		"/././home",
		"/a/b/../home",
	}
	for _, s := range inputs {
		fmt.Printf("%s => %s\n", s, simplifyPath(s))
	}
}

// !FIXME
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

	for _, c := range path {
		if isEmpty(sp) {
			// push the root '/', path guarantees to starts with '/'
			sp = append(sp, c)
		} else if c == '/' {
			// multiple '/' to a single '/'
			if top(sp) != '/' {
				sp = append(sp, '/')
			}
		} else if c == '.' {
			sp = append(sp, '.')
		} else {
			// * handle '.' before push other non-meta characters
			// calc number of '.'
			dots := 0
			for i := len(sp) - 1; i >= 0; i-- {
				if sp[i] != '.' || dots >= 3 {
					break
				}
				dots++
			}

			if dots == 0 || dots >= 3 {
				sp = append(sp, c)
			} else if dots == 1 {
				// pop redundant first before append
				sp[len(sp)-1] = c // override directly
			} else {
				// dots == 2
				if len(sp) == 3 {
					// pop double '.'
					sp = sp[:1]
				} else {
					// pop "/.." fisrt, then keep popping until the first '/'
					sp = sp[:len(sp)-3]
					for i := len(sp) - 1; i >= 0; i-- {
						if sp[i] == '/' {
							break
						}
						sp = sp[:len(sp)-1]
					}

					sp = append(sp, c)
				}
			}
		}
	}

	// TODO: handle trailing '.'

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
