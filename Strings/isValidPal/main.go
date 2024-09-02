package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "A man, a plan, a canal: Panama"

	newStr := []rune(str)
	newStr = toLower(newStr)
	fmt.Println(string(newStr))
	str2 := []rune{}

	for _, ch := range newStr {
		if unicode.IsLetter(ch) {
			str2 = append(str2, ch)
		}
	}
	fmt.Println(string(str2))
	fmt.Println(isValid(string(newStr)))
}

func isValid(str string) bool {

	n := len(str)
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-i-1] {
			return false
		}
	}
	return true
}

func toLower(newStr []rune) []rune {

	for i := 0; i < len(newStr); i++ {
		if newStr[i] >= 'A' && newStr[i] <= 'Z' {
			newStr[i] += 'a' - 'A'
		}
	}
	return newStr
}
