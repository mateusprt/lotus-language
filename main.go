package main

import (
	"github.com/mateusprt/lox-language/vm"
)

func main() {
	virtual_machine := vm.InitVM()

	var constant int = vm.AddConstant(virtual_machine.Chunk, 2)
	vm.WriteChunk(virtual_machine.Chunk, vm.OP_CONSTANT, 123)
	vm.WriteChunk(virtual_machine.Chunk, uint8(constant), 123)

	constant = vm.AddConstant(virtual_machine.Chunk, 2)
	vm.WriteChunk(virtual_machine.Chunk, vm.OP_CONSTANT, 123)
	vm.WriteChunk(virtual_machine.Chunk, uint8(constant), 123)

	vm.WriteChunk(virtual_machine.Chunk, vm.OP_ADD, 123)

	constant = vm.AddConstant(virtual_machine.Chunk, 2)
	vm.WriteChunk(virtual_machine.Chunk, vm.OP_CONSTANT, 123)
	vm.WriteChunk(virtual_machine.Chunk, uint8(constant), 123)

	vm.WriteChunk(virtual_machine.Chunk, vm.OP_DIVIDE, 123)
	vm.WriteChunk(virtual_machine.Chunk, vm.OP_RETURN, 123)

	vm.Interpret(virtual_machine, virtual_machine.Chunk)

	vm.DisassembleChunk(virtual_machine.Chunk, "DEBUG")

}
