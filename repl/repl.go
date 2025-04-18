package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey/lexer"
	"monkey/parser"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	StartAst(in, out)
}

func StartAst(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		// Check for parsing errors
		errors := p.Errors()
		if len(errors) != 0 {
			io.WriteString(out, "Parser errors:\n")
			for _, msg := range errors {
				fmt.Fprintf(out, "\t%s\n", msg)
			}

			continue
		}

		if program != nil {
			for i, stmt := range program.Statements {
				fmt.Fprintf(out, "Statement[%d]: %#v\n", i, stmt)
			}
		}

	}
}

func StartLexer(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
