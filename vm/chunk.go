package vm

const (
	OP_CONSTANT = iota
	OP_ADD
	OP_SUBTRACT
	OP_MULTIPLY
	OP_DIVIDE
	OP_NEGATE
	OP_RETURN
)

type Chunk struct {
	Code      []uint8
	Constants []Value
	lines     []int
}

func WriteChunk(c *Chunk, instruction uint8, line int) {
	c.Code = append(c.Code, instruction)
	c.lines = append(c.lines, line)
}

func AddConstant(c *Chunk, value Value) int {
	c.Constants = append(c.Constants, value)
	return len(c.Constants) - 1
}
