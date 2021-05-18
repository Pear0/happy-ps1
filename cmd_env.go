package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type envPair struct {
	Key   string
	Value string
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Output shell script meant to be eval'd to configure happy-ps1",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("# this output is meant to be interpreted by the current shell.")
		fmt.Println(`if [ -z "$HPS_SKIP" ]; then `)
		for _, pair := range getEnvVars() {
			fmt.Printf("  export %s='%s'\n", pair.Key, pair.Value)
		}
		fmt.Println(`fi`)

	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}

func getEnvVars() []envPair {
	return []envPair{
		{Key: "HPS1", Value: "7518723576123457612"},
		{Key: "HPS_SKIP", Value: "1"},
		{Key: "PS1", Value: fmt.Sprintf("$(%s ps1 --last-exit $?)", os.Args[0])},
	}
}
