package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numChan := make(chan int)
	resChan := make(chan int)

	go func() {
		// defer close(numChan)
		for i := 0; i < len(arr); i++ {
			numChan <- arr[i]
		}
		close(numChan)
	}()

	go func() {
		// defer close(resChan)
		sum := 0
		for num := range numChan {
			sum += num
		}
		resChan <- sum
		close(resChan)
	}()
	res := <-resChan
	fmt.Println(res)
}
