package engine

import (
	"reflect"
	"testing"
)

func TestOp0(t *testing.T) {
	tape := []byte{OP_0}
	engine := NewEngine(tape)

	err := engine.Next()
	if err != nil {
		t.Fatalf("Next failed: %v", err)
	}

	if engine.MainStack.Size() != 1 {
		t.Fatalf("Expected stack size of 1, got %d", engine.MainStack.Size())
	}

	topItem := engine.MainStack.Pop()
	if !reflect.DeepEqual(topItem, []byte{}) {
		t.Errorf("Expected top item to be an empty byte array, got %v", topItem)
	}
}
