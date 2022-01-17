package cmd

import (
	s3 "github.com/LeoFVO/goBucket/pkg/s3Bucket"
	"github.com/spf13/cobra"
)

func scan(cmd *cobra.Command, args []string) {
  wordlist := cmd.Flag("wordlist").Value.String()
  s3.Execute(wordlist)
}

func init() {
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