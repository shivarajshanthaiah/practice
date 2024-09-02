package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	chan1 := make(chan bool)
	chan2 := make(chan bool)

	wg.Add(2)

	go dog(&wg, chan1, chan2)
	go cat(&wg, chan1, chan2)

	chan1 <- true

	wg.Wait()

}

func dog(wg *sync.WaitGroup, chan1, chan2 chan bool) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		<-chan1
		fmt.Println("I am dog")
		chan2 <- true
	}
}

func cat(wg *sync.WaitGroup, chan1, chan2 chan bool) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		<-chan2
		fmt.Println("I am cat")
		if i < 10 {
			chan1 <- true
		}
	}
}
