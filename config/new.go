package config

import (
	"os"

	"github.com/fatih/color"

	_ "embed"
)

//go:embed resources/config.tmpl
var cfg string

func NewConfig() {
	_, err := os.Stat("letc.toml")
	if !os.IsNotExist(err) {
		color.Magenta("目录下已存在config文件")
		return
	}
	file, err := os.Create("letc.toml")
	checkErr(err)
	file.WriteString(cfg)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
