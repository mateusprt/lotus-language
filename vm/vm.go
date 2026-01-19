package vm

import "fmt"

type InterpretResult int

const (
	INTERPRET_OK InterpretResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

type VM struct {
	Chunk              *Chunk
	InstructionPointer int
}

func Interpret(chunk *Chunk) InterpretResult {
	vm := initVM(chunk)
	return run(vm)
}

func initVM(chunk *Chunk) *VM {
	return &VM{Chunk: chunk, InstructionPointer: -1}
}

func run(vm *VM) InterpretResult {
	for {
		readByte(vm)
		switch vm.Chunk.Code[vm.InstructionPointer] {
		case OP_CONSTANT:
			var constant Value = vm.Chunk.Constants[vm.InstructionPointer]
			fmt.Println(constant)
		case OP_RETURN:
			return INTERPRET_OK
		default:
			return INTERPRET_RUNTIME_ERROR
		}
	}
}

func readByte(vm *VM) {
	vm.InstructionPointer++
}
