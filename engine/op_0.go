package engine

// op0 implements the OP_0 opcode of the Bitcoin scripting language.
// OP_0 is a special opcode used to push an empty byte array onto the stack.
// In Bitcoin script, OP_0 is denoted by the byte 0x00. It's often used to represent a false value
// or to provide an initial stack element for scripts, such as in the standard scriptSig for
// a Pay-to-Public-Key-Hash (P2PKH) transaction. It's worth noting that OP_0 is distinct from
// the opcode for pushing the number zero (which would be a single byte array [0]).
// In the context of Bitcoin's script, the OP_0 opcode is important for certain script types
// and is particularly relevant in the construction of multi-signature and SegWit transactions.
func (e *Engine) op0() error {
	e.MainStack.Push([]byte{}) // Push an empty byte array onto the stack
	return nil
}
