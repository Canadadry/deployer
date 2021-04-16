package main

import (
	"app/pkg/monkey/repl"
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) <= 1 {
		return fmt.Errorf("Expect a file to run")
	}
	f, err := os.Open(args[1])
	if err != nil {
		return err
	}
	return repl.Start(f)
}
