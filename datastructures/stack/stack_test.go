package stack

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Len() != 3 {
		t.Error(fmt.Sprintf("Invalid stack length. \"%v\" expected \"%v\" given.", 3, s.Len()))
	}
}

func TestPop(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	i := s.Pop()
	i = s.Pop()

	if i != 2 {
		t.Error(fmt.Sprintf("Invalid stack item. \"%v\" expected \"%v\" given.", 2, i))
	}
}
