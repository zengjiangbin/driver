package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "service name",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("服务: %s", srv.Name())
	},
}

func init() {
	rootCmd.AddCommand(nameCmd)
}
