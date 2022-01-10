package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"syscall"
)

var rootCmd = &cobra.Command{
	Use:   "happy-ps1",
	Short: "Launch a new $SHELL with happy-ps1 integration",
	Run: func(cmd *cobra.Command, args []string) {

		shell := os.Getenv("SHELL")
		if shell == "" {
			shell = "/bin/sh"
		}

		fmt.Printf("running new shell: %s\n", shell)
		fmt.Printf("own location: %s\n\n", os.Args[0])

		var newEnvs []string

		for _, kv := range os.Environ() {
			if !strings.HasPrefix(kv, "PS1=") {
				newEnvs = append(newEnvs, kv)
			}
		}

		for _, pair := range getEnvVars() {
			newEnvs = append(newEnvs, fmt.Sprintf("%s=%s", pair.Key, pair.Value))
		}

		err := syscall.Exec(shell, []string{shell, "-i", "-f", "-o", "PROMPT_SUBST"}, newEnvs)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			os.Exit(1)
		}

	},
}
