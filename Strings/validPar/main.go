package main

import "fmt"

func main() {
	par := "([{{[[[(([]))]]]}}]))))"

	fmt.Println(isValid(par))
}

func isValid(par string) bool {
	match := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}
	for _, ch := range par {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		}

		if ch == ')' || ch == '}' || ch == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != match[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}