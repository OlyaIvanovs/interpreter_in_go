package lexer

import (
	"bufio"
	"fmt"
	"io"

	"github.com/OlyaIvanovs/interpreter_in_go/token"
	"github.com/OlyaIvanovs/interpreter_in_go/lexer"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
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