package main

import (
	"github.com/mateusprt/lox-language/vm"
)

func main() {
	var chunk vm.Chunk
	vm.WriteChunk(&chunk, vm.OP_RETURN)
	vm.AddConstant(&chunk, 42)
	vm.WriteChunk(&chunk, vm.OP_CONSTANT)
	vm.DisassembleChunk(&chunk, "DEBUG")
}
