package cmd

import (
	"github.com/spf13/cobra"
)

func dump(cmd *cobra.Command, args []string) {
  // outputPath := cmd.Flag("outputPath").Value.String()
  // .Execute(outputPath)
}

func init() {
	dumpCmd.Flags().StringP("outputPath", "o", "./output/", "Path of the output folder")
  rootCmd.AddCommand(dumpCmd)
}

var dumpCmd = &cobra.Command{
  Use:   "dump",
  Short: "Dump bucket contents.",
  Long:  `Download all the available content of a bucket if bucket is publicly accessible.`,
  Run: dump,
}