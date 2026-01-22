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
	for {
		readByte(vm)
		switch vm.Chunk.Code[vm.InstructionPointer] {
		case OP_CONSTANT:
			vm.Stack.Push(readConstant(vm))
		case OP_ADD:
			binaryOP(vm, "+")
		case OP_SUBTRACT:
			binaryOP(vm, "-")
		case OP_MULTIPLY:
			binaryOP(vm, "*")
		case OP_DIVIDE:
			binaryOP(vm, "/")
		case OP_NEGATE:
			value, _ := vm.Stack.Pop()
			vm.Stack.Push(-value)
		case OP_RETURN:
			lastValueOnStack, _ := vm.Stack.Pop()
			fmt.Println(lastValueOnStack)
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

func binaryOP(vm *VM, operator string) {
	b, _ := vm.Stack.Pop()
	a, _ := vm.Stack.Pop()
	var result Value
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}
	vm.Stack.Push(result)
}
