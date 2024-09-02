package main

import "fmt"

func main() {
	str1 := "shivaraj"
	str2 := "sha"

	fmt.Println(isSubstring(str1, str2))
}

// func isSubstring(str1, str2 string) bool {
// 	if len(str2) > len(str1) {
// 		return false
// 	}

// 	j := 0
// 	for i := 0; i < len(str1) && j < len(str2); i++ {
// 		if str1[i] == str2[j] {
// 			j++
// 		}
// 	}
// 	return j == len(str2)
// }

func isSubstring(str1, str2 string) bool {
	if len(str2) > len(str1) {
		return false
	}

	for i := 0; i < len(str1); i++ {
		isMatch := true
		for j := 0; j < len(str2); j++ {
			if str1[i+j] != str2[j] {
				isMatch = false
				break
			}
		}
		if isMatch {
			return true
		}
	}
	return false
}
