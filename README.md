# jaryan

A **concurrent-safe FIFO (First-In-First-Out) queue** implementation in Go, inspired by the concept of _flow_ (جریان).

## Features

- Generic support (Go 1.18+)
- Concurrent-safe (`sync.Mutex` + `atomic`)
- Two-lock design (`headMu` for dequeues, `tailMu` for enqueues)
- Provides `Enqueue`, `Dequeue`, `Peek`, and `Len` functions

## Installation

```bash
go get github.com/pooulad/jaryan
```

## Usage

```go
package main

import (
	"fmt"
	"jaryan"
)

func main() {
	q := &jaryan.Queue[int]{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println("Length:", q.Len()) // 3

	if val, ok := q.Peek(); ok {
		fmt.Println("Peek:", val) // 1
	}

	if val, ok := q.Dequeue(); ok {
		fmt.Println("Dequeued:", val) // 1
	}

	fmt.Println("Length after dequeue:", q.Len()) // 2
}
```
## Concurrency Example

```go
var wg sync.WaitGroup
for i := 0; i < 100; i++ {
	wg.Add(1)
	go func(n int) {
		defer wg.Done()
		q.Enqueue(n)
	}(i)
}
wg.Wait()

fmt.Println("Final length:", q.Len())
```

## Benchmarks

Run with:

```bash
go test -bench=. -benchmem
```

```bash
BenchmarkEnqueue-4              25414090                53.36 ns/op           16 B/op          1 allocs/op
BenchmarkDequeue-4              85122091                14.58 ns/op            0 B/op          0 allocs/op
BenchmarkQueueConcurrent-4      14816754               115.7 ns/op            16 B/op          1 allocs/op
```

## License

MIT