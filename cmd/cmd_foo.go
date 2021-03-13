package main

import (
	"github.com/Pear0/hps1/integrations"
	"github.com/spf13/cobra"
)

var fooCmd = &cobra.Command{
	Use:   "foo",
	Short: "Output shell script meant to be eval'd to configure happy-ps1",
	Run: func(cmd *cobra.Command, args []string) {

		integrations.GetMountInfo()

	},
}

func init() {
	rootCmd.AddCommand(fooCmd)
}
