package main

import "fmt"

func main() {
	s := "abccccddggggg"

	fmt.Println(longestPal(s))
}

func longestPal(s string) int {
	freq := make(map[rune]bool)
	res := 0

	for _, c := range s {
		if freq[c] {
			delete(freq, c)
			res += 2
		} else {
			freq[c] = true
		}
	}
	if len(freq) > 0 {
		return res + 1
	}
	return res
}