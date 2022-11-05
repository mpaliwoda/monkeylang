package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mpaliwoda/interpreter-book/lexer"
	"github.com/mpaliwoda/interpreter-book/token"
)

const PROMPT = "$> "

func Start(in io.Reader, out io.Writer) {
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
			if tok.Literal == "exit" {
				fmt.Printf("Bye!\n")
				os.Exit(0)
			}

			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
