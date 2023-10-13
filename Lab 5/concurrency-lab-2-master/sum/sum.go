package main

import (
	"fmt"
	"sync"
)

func main() {
	var sum int
	var wg sync.WaitGroup
	//var lock sync.Mutex
	channel := make(chan int, 1)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//lock.Lock()
			sum = <-channel + 1
			//lock.Unlock()
			channel <- sum
			wg.Done()
		}()
	}

	channel <- 0

	wg.Wait()
	sum = <-channel
	fmt.Println(sum)
}
