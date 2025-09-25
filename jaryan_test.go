package jaryan

import (
	"sync"
	"testing"
)

func TestQueue(t *testing.T) {
	// create queue first.
	q := &Queue[int]{}

	// add orders to it.
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// check length of queue.
	if q.Len() != 3 {
		t.Errorf("expected length 3, got %d", q.Len())
	}

	// remove order from heap.
	val, ok := q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("expected 1, got %v", val)
	}

	// show current heap order
	val, ok = q.Peek()
	if !ok || val != 2 {
		t.Errorf("expected 2 at front, got %v", val)
	}

	val2, ok := q.Dequeue()
	if !ok || val2 != 2 {
		t.Errorf("expected 2, got %v", val2)
	}

	val3, ok := q.Dequeue()
	if !ok || val3 != 3 {
		t.Errorf("expected 3, got %v", val3)
	}

	// now queue is empty. so val4 should be 0 and ok should be false.
	val4, ok := q.Dequeue()
	if ok || val4 != 0 {
		t.Errorf("expected 0, got %v", val4)
	}
}

func TestQueueWithSlice(t *testing.T) {
	// create queue first.
	q := &Queue[int]{}

	// create slice
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// range over slice to add order to queue
	for _, item := range list {
		q.Enqueue(item)
	}

	// check length of queue.
	if q.Len() != len(list) {
		t.Errorf("expected length %d, got %d", len(list), q.Len())
	}

	for i, item := range list {
		// remove order from heap.
		val, ok := q.Dequeue()
		if !ok || val != item {
			t.Errorf("expected %v, got %v", item, val)
		}

		// show current heap order
		if q.Len() > 0 {
			val, ok = q.Peek()
			if !ok || val != list[i+1] {
				t.Errorf("expected %v at front, got %v", list[i+1], val)
			}
		}
	}
}

func TestQueueConcurrent(t *testing.T) {
	q := &Queue[int]{}

	var wg sync.WaitGroup

	// 100 goroutines enqueueing
	for i := range 100 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			q.Enqueue(i)
		}(i)
	}

	wg.Wait()

	if q.Len() != 100 {
		t.Errorf("expected length 100, got %d", q.Len())
	}

	// 100 goroutines dequeueing
	wg = sync.WaitGroup{}
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			q.Dequeue()
		}()
	}

	wg.Wait()

	if q.Len() != 0 {
		t.Errorf("expected empty queue, got length %d", q.Len())
	}
}

// Benchmark enqueue operation
func BenchmarkEnqueue(b *testing.B) {
	q := &Queue[int]{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

// Benchmark dequeue operation
func BenchmarkDequeue(b *testing.B) {
	q := &Queue[int]{}

	// Pre-fill the queue
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}

	b.ResetTimer() // start measuring time

	// dequeue all
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}

// Benchmark concurrent enqueue/dequeue
func BenchmarkQueueConcurrent(b *testing.B) {
	q := &Queue[int]{}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			q.Enqueue(i)
			q.Dequeue()
			i++
		}
	})
}
