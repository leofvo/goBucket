package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version of goBucket",
  Long:  `All software has versions. This is goBucket's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("GoBuster v1.0")
  },
}