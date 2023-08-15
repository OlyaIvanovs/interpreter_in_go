package lexer

import (
	"bufio"
	"fmt"
	"io"

	"github.com/OlyaIvanovs/interpreter_in_go/lexer"
	"github.com/OlyaIvanovs/interpreter_in_go/parser"
	"github.com/OlyaIvanovs/interpreter_in_go/evaluator"
	"github.com/OlyaIvanovs/interpreter_in_go/object"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	
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
		
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, fmt.Sprintf("\t%\n", msg))
	}
}