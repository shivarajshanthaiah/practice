package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var wg sync.WaitGroup
	n := len(arr)
	for i := 0; i < n/2; i++ {
		wg.Add(1)
		go func(i int) {
			wg.Done()
			arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
		}(i)
	}
	wg.Wait()
	fmt.Println(arr)
}
