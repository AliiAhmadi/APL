package main

import (
	"Ahmadi/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
