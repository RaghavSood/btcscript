package engine

import (
	"reflect"
	"testing"
)

func TestOpDup(t *testing.T) {
	engine := NewEngine(nil) // Initializing engine without a script

	item := []byte{1, 2, 3}
	engine.MainStack.Push(item)

	err := engine.opDup()
	if err != nil {
		t.Fatalf("opDup failed: %v", err)
	}

	if engine.MainStack.Size() != 2 {
		t.Fatalf("Expected stack size of 2 after OP_DUP, got %d", engine.MainStack.Size())
	}

	if !reflect.DeepEqual(engine.MainStack.Pop(), item) || !reflect.DeepEqual(engine.MainStack.Pop(), item) {
		t.Fatal("Top two items on the stack should be equal after OP_DUP")
	}
}
