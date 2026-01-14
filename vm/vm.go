package vm

const (
	OP_RETURN = iota
)

type Chunk struct {
	Code []uint8
}

func WriteChunk(c *Chunk, instruction uint8) {
	c.Code = append(c.Code, instruction)
}
