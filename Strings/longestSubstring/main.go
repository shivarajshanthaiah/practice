package main

import (
	"fmt"
)

func main() {
	str := "imaginethepressureofbeinghim"
	fmt.Println(length(str))
}

func length(str string) int {
	res := 0
	for i := 0; i < len(str); i++ {
		unq := make(map[byte]bool)
		currLen := 0
		for j := i; j < len(str); j++ {
			if unq[str[j]] {
				break
			}
			unq[str[j]] = true
			currLen++

		}
		if currLen > res {
			res = currLen
		}
	}
	return res
}
