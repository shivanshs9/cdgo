package cmd

import (
	"os"

	"github.com/shivanshs9/cdgo/pkg/app"
	"github.com/shivanshs9/cdgo/pkg/daemon"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	isBlocking bool
)

func daemonStartHandler(cmd *cobra.Command, args []string) {
	cfg, err := app.InitConfig()
	if err != nil {
		log.Error(err)
		return
	}
	myApp, err := app.InitApp(cfg)
	if err != nil {
		log.Error(err)
		return
	}
	result, err := daemon.Service.StartDaemon(myApp, isBlocking)
	log.Info(result)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func daemonInstallHandler(cmd *cobra.Command, args []string) {
	result, err := daemon.Service.Install("daemon", "--block")
	log.Info(result)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func daemonRemoveHandler(cmd *cobra.Command, args []string) {
	result, err := daemon.Service.Remove()
	log.Info(result)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func daemonStopHandler(cmd *cobra.Command, args []string) {
	result, err := daemon.Service.Stop()
	log.Info(result)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run the daemon in background.",
	Run:   daemonStartHandler,
}

var daemonInstallCmd = &cobra.Command{
	Use: "install",
	Run: daemonInstallHandler,
}

var daemonRemoveCmd = &cobra.Command{
	Use: "remove",
	Run: daemonRemoveHandler,
}

var daemonStopCmd = &cobra.Command{
	Use: "stop",
	Run: daemonStopHandler,
}

func init() {
	daemonCmd.AddCommand(daemonInstallCmd, daemonRemoveCmd, daemonStopCmd)

	daemonCmd.Flags().BoolVarP(&isBlocking, "block", "b", false, "Run daemon as a blocking process")
	rootCmd.AddCommand(daemonCmd)
}
