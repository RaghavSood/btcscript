package engine

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)

	if s.Size() != 2 {
		t.Errorf("Expected stack size of 2, got %d", s.Size())
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)

	popped := s.Pop()
	if popped != 2 {
		t.Errorf("Expected popped value of 2, got %v", popped)
	}

	if s.Size() != 1 {
		t.Errorf("Expected stack size of 1 after pop, got %d", s.Size())
	}
}

func TestStackPopEmpty(t *testing.T) {
	s := NewStack()

	if popped := s.Pop(); popped != nil {
		t.Errorf("Expected nil from popping empty stack, got %v", popped)
	}
}

func TestStackSize(t *testing.T) {
	s := NewStack()

	if s.Size() != 0 {
		t.Errorf("Expected initial stack size of 0, got %d", s.Size())
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("Expected stack size of 3, got %d", s.Size())
	}
}
