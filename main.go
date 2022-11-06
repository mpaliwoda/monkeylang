package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/mpaliwoda/monkeylang/repl"
)

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}
	fmt.Print(MONKEY_FACE)
	fmt.Printf("\nHello %s! This is the Monkey programming language.\n", user.Username)
	fmt.Printf("Feel free to type in commands! To exit, type in: .exit\n\n")
	repl.Start(os.Stdin, os.Stdout)
}
