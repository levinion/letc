package cmd

import (
	"fmt"
	"strings"

	"github.com/levinion/letc/config"
	"github.com/levinion/letc/query"
	"github.com/spf13/cobra"
)

var questionCmd=&cobra.Command{
	Use: "que",
	Aliases: []string{"q"},
	Run: func(cmd *cobra.Command, args []string) {
		config.ReadConfig()
		prompt:=strings.Join(args," ")
		res:=query.GetHelp(prompt)
		fmt.Println(res)
	},
}

func init(){
	rootCmd.AddCommand(questionCmd)
}