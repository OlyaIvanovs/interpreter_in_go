package lexer

import (
	"bufio"
	"fmt"
	"io"

	"github.com/OlyaIvanovs/interpreter_in_go/lexer"
	"github.com/OlyaIvanovs/interpreter_in_go/parser"
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
		p := parser.New(l)
		
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}
		
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, fmt.Sprintf("\t%\n", msg))
	}
}