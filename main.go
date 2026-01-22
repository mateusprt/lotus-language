package main

import (
	"github.com/mateusprt/lox-language/vm"
)

func main() {
	var chunk vm.Chunk

	var constant int = vm.AddConstant(&chunk, 2)
	vm.WriteChunk(&chunk, vm.OP_CONSTANT, 123)
	vm.WriteChunk(&chunk, uint8(constant), 123)

	constant = vm.AddConstant(&chunk, 2)
	vm.WriteChunk(&chunk, vm.OP_CONSTANT, 123)
	vm.WriteChunk(&chunk, uint8(constant), 123)

	vm.WriteChunk(&chunk, vm.OP_ADD, 123)

	constant = vm.AddConstant(&chunk, 2)
	vm.WriteChunk(&chunk, vm.OP_CONSTANT, 123)
	vm.WriteChunk(&chunk, uint8(constant), 123)

	vm.WriteChunk(&chunk, vm.OP_DIVIDE, 123)
	vm.WriteChunk(&chunk, vm.OP_RETURN, 123)

	vm.Interpret(&chunk)

	vm.DisassembleChunk(&chunk, "DEBUG")
}
