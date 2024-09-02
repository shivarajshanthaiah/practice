package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	intChan := make(chan bool)
	charChan := make(chan bool)

	wg.Add(2)

	go func() {
		wg.Done()
		for i := 1; i <= 5; i++ {
			<-intChan
			fmt.Println(i)
			charChan <- true
		}
	}()

	go func() {
		wg.Done()
		for i := 'a'; i <= 'e'; i++ {
			<-charChan
			fmt.Println(string(i))
			intChan <- true
		}
	}()

	intChan <- true
	wg.Wait()
}
