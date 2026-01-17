package vm

const (
	OP_CONSTANT = iota
	OP_RETURN
)

type Chunk struct {
	Code      []uint8
	Constants []Value
	lines     []int
}

func WriteChunk(c *Chunk, instruction uint8) {
	c.Code = append(c.Code, instruction)
}

func AddConstant(c *Chunk, value Value) int {
	c.Constants = append(c.Constants, value)
	return len(c.Constants) - 1
}
