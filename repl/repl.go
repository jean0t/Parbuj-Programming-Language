package repl

import (
	"bufio"
	"fmt"
	"io"
	"tarbuj/lexer"
	"tarbuj/token"
)

const PROMPT = "IM>> "

func Start(in io.Reader, out io.Writer) {
	var scanner = bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		var scanned = scanner.Scan()
		if !scanned {
			return
		}

		var line = scanner.Text()
		var l = lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

}

