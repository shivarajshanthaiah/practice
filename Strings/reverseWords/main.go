package main

import "fmt"

func main() {
	sentence := "The quick brown fox jumped over a lazy dog"
	rev := reversed(sentence)
	fmt.Println(rev)
}

func reversed(s string) string {
	words := []string{}

	idx := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if idx != i {
				words = append(words, s[idx:i])
			}
			idx = i + 1
		}
	}

	if idx < len(s) {
		words = append(words, s[idx:])
	}

	n := len(words)
	for i := 0; i <= n/2; i++ {
		words[i], words[n-i-1] = words[n-i-1], words[i]
	}

	res := ""
	for i := 0; i < len(words); i++ {
		if i != 0 {
			res += " "
		}
		res += words[i]
	}
	return res
}
