package main

import (
	"fmt"
	"sync"
)

func main() {
	odd := make(chan int)
	even := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for num := range odd {
			fmt.Println(num)
		}
	}()

	go func() {
		defer wg.Done()
		for num := range even {
			fmt.Println(num)
		}
	}()

	order := make(chan int)

	go func() {
		defer close(order)
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				order <- i
			} else {
				order <- -i
			}
		}
	}()

	for num := range order {
		if num > 0 {
			even <- num
		} else {
			odd <- -num
		}
	}

	close(odd)
	close(even)

	wg.Wait()
}
