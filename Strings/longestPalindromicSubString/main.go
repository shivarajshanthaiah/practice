package main

import "fmt"

func main() {
	str := "babadracecarmalayalam"
	fmt.Println(longestPal(str))
}

func longestPal(s string) string {
	res := ""
	maxLen := 0

	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > maxLen {
				res = s[l : r+1]
				maxLen = r - l + 1
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > maxLen {
				res = s[l : r+1]
				maxLen = r - l + 1
			}
			l--
			r--
		}
	}
	return res
}
