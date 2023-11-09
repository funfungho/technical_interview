package main

import "fmt"

/*

Category: Stack

Parentheses are mandatory in this problem, even if they are not necessary.

Dijkstra’s Two-Stack Algorithm for Expression Evaluation

The following is not for this problem: https://www.acwing.com/problem/content/3305/

*/

func main() {
	operators := []int32{}
	operands := []int{}

	var s string
	fmt.Scanf("%s", &s)

	numChar := []int32{}
	for _, c := range s {
		switch c {
		// leave out '(' and ' '
		case '(', ' ':
			// do nothing
		case '+', '-', '*', '/':
			operators = append(operators, c)

			if len(numChar) > 0 {
				num := 0
				for _, c := range numChar {
					num = num*10 + int(c-'0')
				}
				operands = append(operands, num)
				numChar = []int32{}
			}
		case ')':
			if len(numChar) > 0 {
				num := 0
				for _, c := range numChar {
					num = num*10 + int(c-'0')
				}
				operands = append(operands, num)
				numChar = []int32{}
			}

			fmt.Println("before", operators, operands)

			op := operators[len(operators)-1]
			operators = operators[:len(operators)-1]

			valRight := operands[len(operands)-1]
			valLeft := operands[len(operands)-2]
			operands = operands[:len(operands)-2]

			switch op {
			case '+':
				operands = append(operands, valLeft+valRight)
			case '-':
				operands = append(operands, valLeft-valRight)
			case '*':
				operands = append(operands, valLeft*valRight)
			case '/':
				operands = append(operands, valLeft/valRight)
			}

			fmt.Println("after", operators, operands)
		default:
			numChar = append(numChar, c)
		}
	}
	fmt.Println(operands[0])
}

/*

((2+3)*(4+5))
(1+((2+3)*(4*5)))

*/
