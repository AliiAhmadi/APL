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
	fmt.Fprint(out, color.Blue("Ahmadi programming language - Copyright (c) 2023 Ali Ahmadi\n\n"))

	for {
		fmt.Fprint(out, color.Yellow(PROMPT))
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
			exit(out)
			os.Exit(0)
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			if evaluated.Type() != object.ERROR_OBJ {
				io.WriteString(out, color.Green(evaluated.Inspect()))
				io.WriteString(out, "\n")
			} else {
				io.WriteString(out, color.Red(evaluated.Inspect()))
				io.WriteString(out, "\n")
			}
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, color.Red("Woops!\n"))
	io.WriteString(out, color.Red(" parser errors:\n"))
	for _, message := range errors {
		io.WriteString(out, color.Red("\t"+message+"\n"))
	}
	io.WriteString(out, "\n")
}

func exit(out io.Writer) {
	fmt.Fprint(out, color.White("goodbye :)\n"))
}
