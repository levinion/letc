package main

import (
	"github.com/levinion/letc/config"

	"github.com/levinion/letc/cmd"
)

func main() {
	config.InitConfig()
	cmd.Execute()
}
