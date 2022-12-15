package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mpaliwoda/monkeylang/evaluator"
	"github.com/mpaliwoda/monkeylang/lexer"
	"github.com/mpaliwoda/monkeylang/object"
	"github.com/mpaliwoda/monkeylang/parser"
)

const PROMPT = "$> "


func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		if line == ".exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil && evaluated != evaluator.NULL {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
