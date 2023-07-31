module github.com/OlyaIvanovs/interpreter_in_go/parser

go 1.19

replace github.com/OlyaIvanovs/interpreter_in_go/ast => ../ast

replace github.com/OlyaIvanovs/interpreter_in_go/token => ../token

replace github.com/OlyaIvanovs/interpreter_in_go/lexer => ../lexer

require (
	github.com/OlyaIvanovs/interpreter_in_go/ast v0.0.0-00010101000000-000000000000
	github.com/OlyaIvanovs/interpreter_in_go/lexer v0.0.0-20230730095935-78656e90f586
	github.com/OlyaIvanovs/interpreter_in_go/token v0.0.0-20230730095935-78656e90f586
)
