package queue

import "sync"

type Item interface{}

type Queue struct {
	items []Item
	mu    sync.Mutex
}

func (q *Queue) Push(i Item) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.items = append(q.items, i)
}

func (q *Queue) Peek() Item {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil
	}

	return q.items[0]
}

func (q *Queue) Pop() Item {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil
	}

	i := q.items[0]
	q.items = q.items[1:]

	return i
}

func (q *Queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.items)
}

func New() *Queue {
	return &Queue{}
}
