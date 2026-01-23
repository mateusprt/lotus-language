package main

import (
	"fmt"
	"os"

	"github.com/mateusprt/lox-language/cmd"
)

func main() {

	if len(os.Args) == 1 {
		cmd.Repl()
		return
	}

	if len(os.Args) == 2 {
		cmd.RunFile(os.Args[1])
		return
	}

	fmt.Println("Usage: lotues [path]")
	os.Exit(64)
}
