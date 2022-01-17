package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	runCmd.Flags().StringP("wordlist", "w", "", "Path of the wordlist to use")
  rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
  Use:   "run",
  Short: "Run all commands on bucket(s)",
  Long:  `Run a complete scan and exploit on buckets.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("")
  },
}