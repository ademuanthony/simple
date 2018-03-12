package main

import (
	"os/user"
	"fmt"
	"os"
	"github.com/ademuanthony/simple/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Print(repl.SIMPLE)
	fmt.Printf("\nHello %s! This is the SIMPLE programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
