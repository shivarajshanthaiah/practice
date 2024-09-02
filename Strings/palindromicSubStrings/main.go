package main

import "fmt"

func main() {
	str := "tlthpowwythupxaszmxhqbfbxegiqz"

	fmt.Println(countSubString(str))
}

func countSubString(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		ans += isPal(s, i, i)
		ans += isPal(s, i, i+1)
	}
	return ans
}

func isPal(s string, l, r int) int {
	count := 0

	for l >= 0 && r < len(s) && s[l] == s[r] {
		count++
		l--
		r++
	}
	return count
}
