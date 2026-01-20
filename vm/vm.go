package vm

import (
	"fmt"

	"github.com/mateusprt/lox-language/utils"
	"github.com/mateusprt/lox-language/utils/data_structures"
)

type InterpretResult int

const (
	INTERPRET_OK InterpretResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

const STACK_MAX = 256

type VM struct {
	Chunk              *Chunk
	InstructionPointer int
	Stack              *data_structures.Stack[Value]
	StackTop           Value
}

func Interpret(chunk *Chunk) InterpretResult {
	vm := initVM(chunk)
	return run(vm)
}

func initVM(chunk *Chunk) *VM {
	return &VM{
		Chunk:              chunk,
		InstructionPointer: -1,
		Stack:              utils.NewStack[Value](STACK_MAX),
	}
}

func run(vm *VM) InterpretResult {
	fmt.Println("          ")
	for slot := 0; slot < len(vm.Stack.Elements); slot++ {
		fmt.Printf("[ ")
		fmt.Printf("%v", vm.Stack.Elements[slot])
		fmt.Printf(" ]")
	}
	DisassembleChunk(vm.Chunk, "code")
	for {
		readByte(vm)
		switch vm.Chunk.Code[vm.InstructionPointer] {
		case OP_CONSTANT:
			var constant Value = readConstant(vm)
			vm.Stack.Push(constant)
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

func readConstant(vm *VM) Value {
	return vm.Chunk.Constants[vm.InstructionPointer]
}
