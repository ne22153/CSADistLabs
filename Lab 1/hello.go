package main

import (
	"fmt"
	"time"
)

func printer(i int) {
	fmt.Println("Hello from goroutine ", i)
}

func main() {
	for i := 0; i < 5; i++ {
		go printer(i)
	}
	time.Sleep(1 * time.Millisecond)
}
