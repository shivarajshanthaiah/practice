package main

import (
	"fmt"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(2)

	ch := make(chan int)

	arr1 := []int{100, 20, 30, 40, 50}
	arr2 := []int{10, 20, 300, 40, 50}

	go func() {
		// wg.Done()
		sum := 0
		for _, val := range arr1 {
			sum += val
		}
		ch <- sum
	}()

	go func() {
		// wg.Done()
		sum := 0
		for _, val := range arr2 {
			sum += val
		}
		ch <- sum
	}()

	res1 := <-ch
	res2 := <-ch

	// wg.Wait()

	if res1 > res2 {
		fmt.Printf("arr1 is greater : %d", res1)
	} else {
		fmt.Printf("arr2 is greater : %d", res2)
	}

	fmt.Println("")

}
