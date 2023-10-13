package main

import (
	"fmt"
	"github.com/ChrisGora/semaphore"
	"math/rand"
	"time"
)

type buffer struct {
	b                 []int
	size, read, write int
}

func newBuffer(size int) buffer {
	return buffer{
		b:     make([]int, size),
		size:  size,
		read:  0,
		write: 0,
	}
}

func (buffer *buffer) get() int {
	x := buffer.b[buffer.read]
	fmt.Println("Get\t", x, "\t", buffer)
	buffer.read = (buffer.read + 1) % len(buffer.b)
	return x
}

func (buffer *buffer) put(x int) {
	buffer.b[buffer.write] = x
	fmt.Println("Put\t", x, "\t", buffer)
	buffer.write = (buffer.write + 1) % len(buffer.b)
}

func producer(buffer *buffer, start, delta int, spaceAvailable semaphore.Semaphore, workAvailable semaphore.Semaphore, lock semaphore.Semaphore) {
	x := start
	for {
		spaceAvailable.Wait()
		lock.Wait()
		buffer.put(x)
		x = x + delta
		workAvailable.Post()
		lock.Post()
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func consumer(buffer *buffer, spaceAvailable semaphore.Semaphore, workAvailable semaphore.Semaphore, lock semaphore.Semaphore) {
	for {
		workAvailable.Wait()
		lock.Wait()
		_ = buffer.get()
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		spaceAvailable.Post()
		lock.Post()
	}
}

func main() {
	buffer := newBuffer(5)

	spaceAvailable := semaphore.Init(5, 5)
	workAvailable := semaphore.Init(5, 0)

	lock := semaphore.Init(1, 1)

	go producer(&buffer, 1, 1, spaceAvailable, workAvailable, lock)
	go producer(&buffer, 1000, -1, spaceAvailable, workAvailable, lock)

	consumer(&buffer, spaceAvailable, workAvailable, lock)
}
