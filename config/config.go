package config

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func ReadConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		color.Cyan("Config not exist. Try use `letc init`.")
		os.Exit(1)
	}
}

func InitConfig() {
	viper.SetConfigFile("./letc.toml")
	viper.SetDefault("alias.easy", "easy")
	viper.SetDefault("alias.medium", "medium")
	viper.SetDefault("alias.hard", "hard")
}
