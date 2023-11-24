package repl

import (
	"Ahmadi/lexer"
	"Ahmadi/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT string = "APL>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	fmt.Fprintf(out, "Ahmadi programming language - Copyright (c) 2023 Ali Ahmadi\n\n")

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
		fmt.Fprintf(out, "\n")
	}
}
