package vm

import "fmt"

func DisassembleChunk(c *Chunk, name string) {
	fmt.Printf("== %s ==\n", name)
	for offset := 0; offset < len(c.Code); {
		offset = disassembleInstruction(c, offset)
	}
}

func disassembleInstruction(c *Chunk, offset int) int {
	instruction := c.Code[offset]
	switch instruction {
	case OP_RETURN:
		return simpleInstruction("OP_RETURN", offset)
	case OP_CONSTANT:
		return constantInstruction("OP_CONSTANT", c, offset)
	default:
		fmt.Printf("Unknown opcode %d\n", instruction)
		return offset + 1
	}
}

func simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}

func constantInstruction(name string, chunk *Chunk, offset int) int {
	var constant uint8 = chunk.Code[offset]
	fmt.Printf("%-16s %4d ", name, constant)
	fmt.Printf("'%v'\n", chunk.Constants[constant])
	return offset + 1
}
