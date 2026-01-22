package main

import (
	"github.com/mateusprt/lox-language/vm"
)

func main() {
	var chunk vm.Chunk

	vm.WriteChunk(&chunk, vm.OP_CONSTANT, 0)
	vm.AddConstant(&chunk, 1.2)
	vm.WriteChunk(&chunk, vm.OP_NEGATE, 0)
	vm.WriteChunk(&chunk, vm.OP_RETURN, 0)
	vm.Interpret(&chunk)
	vm.DisassembleChunk(&chunk, "DEBUG")
}
