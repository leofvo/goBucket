package cmd

import (
	s3 "github.com/LeoFVO/goBucket/pkg/s3Bucket"
	"github.com/spf13/cobra"
)

func scan(cmd *cobra.Command, args []string) {
  wordlist := cmd.Flag("wordlist").Value.String()
  onlyCritical, _ := cmd.Flags().GetBool("only-critical")
  s3.Execute(wordlist, onlyCritical)
}

func init() {
  scanCmd.Flags().BoolP("only-critical", "", false, "Should only search for critical files")
	scanCmd.Flags().StringP("wordlist", "w", "", "Path of the wordlist to use")
  scanCmd.MarkFlagRequired("wordlist")
  rootCmd.AddCommand(scanCmd)
}

var scanCmd = &cobra.Command{
  Use:   "scan",
  Short: "Scan bucket contents.",
  Long:  `Scan for all bucket contents and dumps it if bucket is publicly accessible.`,
  Run: scan,
}