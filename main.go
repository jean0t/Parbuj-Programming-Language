package main

import (
	"fmt"
	"os"
	"os/user"
	"tarbuj/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Parbuj interactive mode\n", user.Username)
	fmt.Printf("Version 0.1 - Type any command\n")
	repl.Start(os.Stdin, os.Stdout)
}
