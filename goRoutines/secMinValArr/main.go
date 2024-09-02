package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numsChan := make(chan int)
	resChan := make(chan [2]int)

	go func() {
		defer close(numsChan)
		for _, num := range arr {
			numsChan <- num
		}
	}()

	go func() {
		defer close(resChan)

		fm, sm := math.MaxInt, math.MaxInt
		for num := range numsChan {
			if num < fm {
				sm, fm = fm, num
			} else if num > fm && num < sm {
				sm = num
			}
		}
		resChan <- [2]int{fm, sm}
	}()

	res := <-resChan
	fmt.Println(res[1])
}
