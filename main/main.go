package main

import (
	"fmt"
	"os"
	"github.com/haitammabrouk/repl"
	)

func main() {
	fmt.Printf("Go Ahead Type Some Code !\n")
	repl.Start(os.Stdin, os.Stdout)
}