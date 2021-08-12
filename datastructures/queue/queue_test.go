package queue

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Len() != 3 {
		t.Error(fmt.Sprintf("Invalid queue length. \"%v\" expected \"%v\" given.", 3, q.Len()))
	}
}

func TestPop(t *testing.T) {
	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	i := q.Pop()

	if i != 1 {
		t.Error(fmt.Sprintf("Invalid queue item. \"%v\" expected \"%v\" given.", 1, i))
	}
}

func TestPeek(t *testing.T) {
	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	i := q.Peek()

	if i != 1 {
		t.Error(fmt.Sprintf("Invalid queue item. \"%v\" expected \"%v\" given.", 1, i))
	}

	if q.Len() != 3 {
		t.Error(fmt.Sprintf("Invalid queue length. \"%v\" expected \"%v\" given.", 3, q.Len()))
	}
}
