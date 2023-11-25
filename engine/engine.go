package engine

import "errors"

const (
	OP_0     = 0x00
	OP_FALSE = OP_0 // An alias for OP_0
	OP_DUP   = 0x76 // Duplicates the top stack item
)

type Engine struct {
	MainStack *Stack
	AltStack  *Stack
	Tape      []byte
	Pointer   int
}

func NewEngine(tape []byte) *Engine {
	return &Engine{
		MainStack: NewStack(),
		AltStack:  NewStack(),
		Tape:      tape,
		Pointer:   0,
	}
}

func (e *Engine) Next() error {
	if e.Pointer >= len(e.Tape) {
		return errors.New("end of script")
	}

	opcode := e.Tape[e.Pointer]
	e.Pointer++

	switch opcode {
	case OP_0:
		return e.op0()
	case OP_DUP:
		return e.opDup()
	default:
		return errors.New("unknown opcode")
	}
}
