package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zengjiangbin/driver/structs/service"
	"os"
)

var srv service.Service

var rootCmd = &cobra.Command{
	Use: "service loader",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.Get("port"))
		fmt.Println(555)
	},
}

func Execute(s service.Service) {
	srv = s
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Int("port", 0, "service rpc port")
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
}

func loadService() {

}
