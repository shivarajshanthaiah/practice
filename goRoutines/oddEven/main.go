package main

import (
	"fmt"
	"sync"
)

func main() {
	oddCh := make(chan int)
	evenCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go printOdd(&wg, oddCh)
	go printeven(&wg, evenCh)

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			oddCh <- i
		} else {
			evenCh <- i
		}
	}

	close(oddCh)
	close(evenCh)
	wg.Wait()
}

func printOdd(wg *sync.WaitGroup, oddCh chan int) {
	defer wg.Done()
	for num := range oddCh {
		fmt.Println(num)
	}
}

func printeven(wg *sync.WaitGroup, evenCh chan int) {
	defer wg.Done()
	for num := range evenCh {
		fmt.Println(num)
	}
}
