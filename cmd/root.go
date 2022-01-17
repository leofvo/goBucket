package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "goBucket",
		Short: "Powerful bucket scanner, analyser & dumper.",
		Long: `goBucket is an powerful utilitary to exploit cloud-providers buckets. Provide multiple tools such as scan, dumps files, search for creds, determine permissions, etc. 
Support: S3 API.
		
Developped by @LeoFVO.
Do not use to gain non-authorized access. Use it at your own risk.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output (errors)")
	rootCmd.PersistentFlags().IntP("threads", "t", 10, "Number of concurrent threads")
}
