package repl

import (
	"Ahmadi/evaluator"
	"Ahmadi/lexer"
	"Ahmadi/object"
	"Ahmadi/parser"
	"bufio"
	"fmt"
	"io"
	"os"
)

const PROMPT string = "APL>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	fmt.Fprintf(out, "Ahmadi programming language - Copyright (c) 2023 Ali Ahmadi\n\n")

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)
		parser := parser.New(lex)
		program := parser.ParseProgram()
		if len(parser.Errors()) != 0 {
			printParserErrors(out, parser.Errors())
			continue
		}

		if program.Statements[0].String() == "exit" {
			os.Exit(0)
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops!\n")
	io.WriteString(out, " parser errors:\n")
	for _, message := range errors {
		io.WriteString(out, "\t"+message+"\n")
	}
	io.WriteString(out, "\n")
}
