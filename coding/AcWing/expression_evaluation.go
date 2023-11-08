package main

import "fmt"

// https://www.acwing.com/problem/content/3305/

func main() {
	operators := []int32{}
	operands := []int{}

	var s string
	fmt.Scanf("%s", &s)

	numChar := []int32{}
	for _, c := range s {
		switch c {
		case '(', ')', '+', '-', '*', '/':
			operators = append(operators, c)

			if len(numChar) > 0 {
				num := 0
				for _, c := range numChar {
					num = num*10 + int(c-'0')
				}
				operands = append(operands, num)
				numChar = []int32{}
			}
		default:
			numChar = append(numChar, c)
		}
	}
	fmt.Println(operators, operands)
}

/*

(2+2)*(1+123)

*/
