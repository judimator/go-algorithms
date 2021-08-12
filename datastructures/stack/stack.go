package stack

import "sync"

type Item interface{}

type Stack struct {
	items []Item
	mu    sync.Mutex
}

func (s *Stack) Push(i Item) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items = append(s.items, i)
}

func (s *Stack) Pop() Item {
	s.mu.Lock()
	defer s.mu.Unlock()

	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]

	return item
}

func (s *Stack) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.items)
}

func New() *Stack {
	return &Stack{}
}
