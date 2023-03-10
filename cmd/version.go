package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "service version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("æå¡: %s \nçæ¬: %s", srv.Name(), srv.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
