package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "trrrreereeeewwwess"

	fmt.Println(sorted(str))
}

func sorted(str string) string {
	freq := make(map[rune]int)

	for _, ch := range str {
		freq[ch]++
	}

	con := []rune(str)

	for i := 0; i < len(con); i++ {
		sort.Slice(con, func(i, j int) bool {
			if freq[con[i]] == freq[con[j]] {
				return con[i] < con[j]
			}
			return freq[con[i]] > freq[con[j]]
		})
	}
	return string(con)
}

