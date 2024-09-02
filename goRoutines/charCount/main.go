package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	par := "This is a sample paragraph. If there are any characters left in the map, it means at least one character can be placed in the center of the palindrome, allowing you to increase the palindrome length by 1"

	count := make(chan map[rune]int)
	// var wg sync.WaitGroup
	// wg.Add(1)

	go func() {
		// wg.Done()
		freq := make(map[rune]int)
		for _, c := range strings.ToLower(par) {
			if unicode.IsLetter(c) {
				freq[c]++
			}
		}

		count <- freq
		close(count)
	}()

	// wg.Wait()
	for c := range count {
		for char, count := range c {
			fmt.Printf("%c: %d, ", char, count)
		}
	}
	fmt.Println("")

}
