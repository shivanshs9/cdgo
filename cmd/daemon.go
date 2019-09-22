package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(daemonCmd)
}

func daemonHandler(cmd *cobra.Command, args []string) {

}

var daemonCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the daemon in background.",
	Run:   daemonHandler,
}
