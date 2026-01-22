package vm

import "fmt"

func DisassembleChunk(c *Chunk, name string) {
	fmt.Printf("== %s ==\n", name)
	for offset := 0; offset < len(c.Code); {
		offset = disassembleInstruction(c, offset)
	}
}

func disassembleInstruction(c *Chunk, offset int) int {
	if offset > 0 && c.lines[offset] == c.lines[offset-1] {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%4d ", c.lines[offset])
	}
	instruction := c.Code[offset]
	switch instruction {
	case OP_ADD:
		return simpleInstruction("OP_ADD", offset)
	case OP_SUBTRACT:
		return simpleInstruction("OP_SUBTRACT", offset)
	case OP_MULTIPLY:
		return simpleInstruction("OP_MULTIPLY", offset)
	case OP_DIVIDE:
		return simpleInstruction("OP_DIVIDE", offset)
	case OP_NEGATE:
		return simpleInstruction("OP_NEGATE", offset)
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

func ShowStackTracing(vm *VM) {
	fmt.Println("          ")
	fmt.Println(len(vm.Stack.Elements))
	for slot := 0; slot < len(vm.Stack.Elements); slot++ {
		fmt.Printf("[%v]\n", vm.Stack.Elements[slot])
	}
	DisassembleChunk(vm.Chunk, "code")
}
