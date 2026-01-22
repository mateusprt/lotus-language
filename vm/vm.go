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
}

func Interpret(chunk *Chunk) InterpretResult {
	vm := initVM(chunk)
	return run(vm)
}

func initVM(chunk *Chunk) *VM {
	return &VM{
		Chunk:              chunk,
		InstructionPointer: 0,
		Stack:              utils.NewStack[Value](STACK_MAX),
	}
}

func run(vm *VM) InterpretResult {
	for {
		current_instruction := readByte(vm)
		switch current_instruction {
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

func readByte(vm *VM) uint8 {
	instruction := vm.Chunk.Code[vm.InstructionPointer]
	vm.InstructionPointer++
	return instruction
}

func readConstant(vm *VM) Value {
	constantIndex := readByte(vm)
	return vm.Chunk.Constants[constantIndex]
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
