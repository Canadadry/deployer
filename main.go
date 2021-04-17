package main

import (
	"app/internal"
	"fmt"
	"os"
)

func main() {
	if err := internal.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
