package jaryan

import (
	"sync"
	"sync/atomic"
)

// Node represents a single order in the queue with generic type
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// Queue holds pointers to the head (front) and tail (back) with generic type
type Queue[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	headMu sync.Mutex
	tailMu sync.Mutex
	size   int64
}

// Enqueue adds an order at the back (tail)
func (q *Queue[T]) Enqueue(value T) {
	q.tailMu.Lock()
	defer q.tailMu.Unlock()

	newNode := &Node[T]{Value: value}
	if q.tail != nil {
		q.tail.Next = newNode
	}
	q.tail = newNode
	if q.head == nil {
		q.head = newNode
	}
	q.size++
}

// Dequeue removes an order from the front (head)
func (q *Queue[T]) Dequeue() (T, bool) {
	q.headMu.Lock()
	defer q.headMu.Unlock()

	if q.head == nil {
		var zero T
		return zero, false
	}
	value := q.head.Value
	q.head = q.head.Next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return value, true
}

// Peek just shows the front order without removing
func (q *Queue[T]) Peek() (T, bool) {
	q.headMu.Lock()
	defer q.headMu.Unlock()

	if q.head == nil {
		var zero T
		return zero, false
	}
	return q.head.Value, true
}

// Len returns length of queue
func (q *Queue[T]) Len() int {
	return int(atomic.LoadInt64(&q.size))
}
