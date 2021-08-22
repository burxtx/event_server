package app

import (
	"flag"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "event",
		Short: "event server",
	}
)

func Execute() error {
	rootCmd.Flags().AddGoFlagSet(flag.CommandLine)
	return rootCmd.Execute()
}
