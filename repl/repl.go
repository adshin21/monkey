// Package repl will be responsible for reading, evaluating, printing and looping
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/adshin21/monkey/lexer"
	"github.com/adshin21/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
