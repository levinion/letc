package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "leetcode辅助工具包，功能同 letc cal",
	Use:   "letc",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Execute() {
	rootCmd.Execute()
}
