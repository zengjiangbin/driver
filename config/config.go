package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var C = config{viper.New()}

type config struct {
	*viper.Viper
}

func (c *config) Init(f map[string]*pflag.Flag) {
	c.bindPFlags(f)
	c.readConfig()
	c.mergeGlobalConfig()
}

func (c *config) bindPFlags(f map[string]*pflag.Flag) {
	for k, v := range f {
		if err := c.BindPFlag(k, v); err != nil {
			panic(err)
		}
	}
}

func (c *config) readConfig() {
	c.SetConfigFile(c.GetString("config"))
	if err := c.ReadInConfig(); err != nil {
		panic(err)
	}
}

func (c *config) mergeGlobalConfig() {
	c.SetConfigFile(c.GetString("global_config"))
	_ = c.MergeInConfig()
}
