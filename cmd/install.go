package cmd

import (
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install uptimekuma-cli",
	Long: `Install uptimekuma-cli.
On proxmox, you have the option to install it to all containers.
If compat is installed, you can optionally monitor uptime for all containers too.
`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
