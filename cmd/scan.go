package cmd

import (
	s3 "github.com/LeoFVO/goBucket/pkg/s3"
	"github.com/spf13/cobra"
)

func scan(cmd *cobra.Command, args []string) {
  mode := cmd.Flag("mode").Value.String()

  onlyCritical, _ := cmd.Flags().GetBool("only-critical")
  s3.Execute(mode, args[0], onlyCritical)
}

func init() {
  scanCmd.Flags().BoolP("only-critical", "", false, "Should only search for critical files")
	scanCmd.Flags().StringP("mode", "m", "url", "type of target: url or wordlist")
  rootCmd.AddCommand(scanCmd)
}

var scanCmd = &cobra.Command{
  Use:   "scan [flags] [url/wordlist path]",
  Short: "Scan bucket contents.",
  Long:  `Scan for all bucket contents and write all open url in "scan_result.txt"`,
  Args: cobra.MinimumNArgs(1),
  Run: scan,
}