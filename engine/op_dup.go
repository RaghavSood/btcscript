package engine

import "fmt"

// opDup duplicates the top item on the stack. It is a fundamental
// operation in Bitcoin scripts, used for copying items for subsequent
// operations like checking a signature against a public key.
func (e *Engine) opDup() error {
	if e.MainStack.Size() < 1 {
		return fmt.Errorf("stack underflow: OP_DUP requires an item on the stack")
	}

	topItem := e.MainStack.Pop()
	e.MainStack.Push(topItem)
	e.MainStack.Push(topItem)

	return nil
}
