package main

import (
	"fmt"

	"github.com/mateusprt/lox-language/vm"
)

func main() {
	var chunk vm.Chunk
	vm.WriteChunk(&chunk, vm.OP_RETURN)
	fmt.Println(chunk.Code)
}
