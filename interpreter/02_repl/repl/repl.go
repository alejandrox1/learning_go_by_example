package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/alejandrox1/learning_go_by_example/interpreter/second_repl/lexer"
	"github.com/alejandrox1/learning_go_by_example/interpreter/second_repl/token"
)

const PROMPT = ">>"

func Start(in io.Reader,out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
