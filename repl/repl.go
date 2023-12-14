package repl

import (
	"Ahmadi/color"
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
	fmt.Fprintf(out, color.BLUE+"Ahmadi programming language - Copyright (c) 2023 Ali Ahmadi\n\n"+color.RESET)

	for {
		fmt.Fprint(out, color.YELLOW+PROMPT+color.RESET)
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
			io.WriteString(out, color.GREEN+evaluated.Inspect()+color.RESET)
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, color.RED+"Woops!\n"+color.RESET)
	io.WriteString(out, color.RED+" parser errors:\n"+color.RESET)
	for _, message := range errors {
		io.WriteString(out, color.RED+"\t"+message+"\n"+color.RESET)
	}
	io.WriteString(out, "\n")
}
