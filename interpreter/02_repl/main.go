package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/alejandrox1/learning_go_by_example/interpreter/second_repl/repl"
)


func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the SuperGo programming language!\n", user.Username)
	fmt.Printf("Feel free to type in any commands...\n")
	repl.Start(os.Stdin, os.Stdout)
}
