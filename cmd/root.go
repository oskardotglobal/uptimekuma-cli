package cmd

import (
	"github.com/go-co-op/gocron"
	"github.com/oskardotglobal/uptimekuma-cli/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
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
		_, err := scheduler.Every(1).Minute().Do(util.ReportStatus, viper.GetViper(), "nodes.root")
		util.CheckErrorWithMsg(err, "Error whilst scheduling task")

		// Schedule tasks for nodes running on root node
		//compat.ReportNodes(scheduler, viper.GetViper())

		defer scheduler.StartAsync()
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.uptimekuma-cli.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		util.CheckError(err)

		// Search config in home directory with name ".uptimekuma-cli.yaml"
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".uptimekuma-cli.yaml")

		viper.Set("nodes.root", "your_url_here")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		util.Info("Using config file " + viper.ConfigFileUsed())
	}
}
