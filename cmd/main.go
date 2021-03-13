package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func main() {

	// we always want to output color escape sequences, even if the output is not a tty
	color.NoColor = false

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
