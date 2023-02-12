package cmd

import (
	_ "embed"

	"github.com/levinion/letc/config"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Short: "初始化配置文件",
	Use:   "init",
	Run: func(cmd *cobra.Command, args []string) {
		config.NewConfig()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
