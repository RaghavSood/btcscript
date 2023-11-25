package engine

import "errors"

const (
	OP_0     = 0x00
	OP_FALSE = OP_0 // An alias for OP_0
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
	// Add cases for other opcodes here
	default:
		return errors.New("unknown opcode")
	}
}
