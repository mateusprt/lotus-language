package main

import (
	"github.com/mateusprt/lox-language/vm"
)

func main() {
	var chunk vm.Chunk
	vm.WriteChunk(&chunk, vm.OP_RETURN)
	vm.DisassembleChunk(&chunk, "DEBUG")
}
