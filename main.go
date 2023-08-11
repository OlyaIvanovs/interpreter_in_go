package main

import (
	"fmt"
	"os"
	"os/user"
	
	repl "github.com/OlyaIvanovs/interpreter_in_go/repl"
)

func main() {
	user , err := user.Current()
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Hello %s! This is the New Programming language\n", user.Username)
	fmt.Printf("Type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

