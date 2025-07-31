package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/adshin21/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Monkey here!\n", user.Username)
	fmt.Printf("Give me commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
