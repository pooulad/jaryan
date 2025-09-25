package main

import (
	"fmt"
	"sync"

	"github.com/pooulad/jaryan"
)

func main() {
	// create a queue of integers
	q := &jaryan.Queue[int]{}

	// Enqueue some items
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println("Initial queue length:", q.Len())

	// Peek at the front
	if val, ok := q.Peek(); ok {
		fmt.Println("Peek front:", val) // should be 10
	}

	// Dequeue items
	if val, ok := q.Dequeue(); ok {
		fmt.Println("Dequeued:", val) // should be 10
	}

	if val, ok := q.Dequeue(); ok {
		fmt.Println("Dequeued:", val) // should be 20
	}

	fmt.Println("Queue length after two dequeues:", q.Len())

	// Demonstrate concurrency
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			q.Enqueue(n * 100)
		}(i)
	}

	wg.Wait()
	fmt.Println("Final queue length (after goroutines):", q.Len())

	// Drain the queue
	for {
		val, ok := q.Dequeue()
		if !ok {
			break
		}
		fmt.Println("Drained:", val)
	}
}
