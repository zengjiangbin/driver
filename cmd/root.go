package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/zengjiangbin/driver/config"
	"github.com/zengjiangbin/driver/structs/service"
	"os"
)

var srv service.Service

var rootCmd = &cobra.Command{
	Use: "service loader",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.C.Get("port"))
		fmt.Println(config.C.Get("hport"))
		fmt.Println(config.C.Get("rpc_server_name"))
	},
}

func Execute(s service.Service) {
	srv = s
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

const (
	configFlag       = "config"
	globalConfigFlag = "global_config"
	portFlag         = "port"
	hportFlag        = "hport"
)

func init() {
	rootCmd.PersistentFlags().String(configFlag, "./config.yml", "service config")
	rootCmd.PersistentFlags().String(globalConfigFlag, "../global.yml", "service global config")
	rootCmd.PersistentFlags().Int(portFlag, 0, "service rpc port")
	rootCmd.PersistentFlags().Int(hportFlag, 0, "service rpc health port")
	cobra.OnInitialize(cobraInit)
}

func cobraInit() {
	f := map[string]*pflag.Flag{
		configFlag:       rootCmd.PersistentFlags().Lookup(configFlag),
		globalConfigFlag: rootCmd.PersistentFlags().Lookup(globalConfigFlag),
		portFlag:         rootCmd.PersistentFlags().Lookup(portFlag),
		hportFlag:        rootCmd.PersistentFlags().Lookup(hportFlag),
	}
	config.C.Init(f)
}
