package main

import "fmt"

func main() {
	str := "AABAAABBSDAAABBBBB"
	k := 3

	fmt.Println(charReplace(str, k))
}

func charReplace(str string, k int) int {
	freq := make([]int, 26)
	count := 0
	res := 0

	for i, j := 0, 0; i < len(str); i++ {
		freq[str[i]-'A']++
		count++

		for count-maxFreq(freq) > k {
			freq[str[j]-'A']--
			count--
			j++
		}
		res = max(res, count)
	}
	return res
}

func maxFreq(str []int) int {
	ans := 0
	for _, v := range str {
		ans = max(ans, v)
	}
	return ans
}
