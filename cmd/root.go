package cmd

import (
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/oskardotglobal/uptimekuma-cli/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	scheduler *gocron.Scheduler
)

var rootCmd = &cobra.Command{
	Use:   "uptimekuma-cli",
	Short: "Ping uptime kuma server every minute",
	Long:  `Cli tool to report uptime to uptime kuma using push method`,
	Run: func(cmd *cobra.Command, args []string) {
		scheduler = gocron.NewScheduler(time.FixedZone("Europe/Berlin", 1*60*60))

		// Schedule task for root node
		_, err := scheduler.Every(1).Minute().Do(util.ReportStatus, "nodes.root")
		util.CheckErrorWithMsg(err, "Error whilst scheduling task")

		// Schedule tasks for nodes running on root node
		//compat.ReportNodes(scheduler, viper.GetViper())

		scheduler.StartBlocking()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	util.InitConfig(viper.GetViper())
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.uptimekuma-cli.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		util.CheckError(err)

		cfgFile = home + "/.uptimekuma-cli.yaml"
	}

	if _, err := os.ReadFile(cfgFile); err != nil {
		_, err := os.Create(cfgFile)
		util.CheckErrorWithMsg(err, "Couldn't create config file!")
	}

	viper.SetConfigFile(cfgFile)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		util.Info("Using config file " + viper.ConfigFileUsed())
	}

    util.SetNodeUrlIfEmpty("nodes.root")
}
