package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version subcommand show stocksUntil version info.",

	Run: func(cmd *cobra.Command, args []string) {
		output, err := ExecuteCommand("git", "rev-parse", "--short", "HEAD")
		if err != nil {
			Error(cmd, args, err)
		}

		_, err = fmt.Fprint(os.Stdout, "Stocks Version 1.0.0_"+output)
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
