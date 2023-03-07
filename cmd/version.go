package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "service version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("服务: %s \n版本: %s", srv.Name(), srv.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
