package cmd

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/levinion/letc/timer"

	"github.com/levinion/letc/config"

	"github.com/levinion/letc/query"

	"github.com/levinion/letc/model"

	"github.com/fatih/color"
	"github.com/lunny/html2md"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "embed"
)

var getCmd = &cobra.Command{
	Short: "通过题号拉取具体题目",
	Use:   "get",
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		checkErr(err)
		Get(index)
	},
}

func Get(index int) {
	t := timer.NewTimer()
	t.Start()
	config.ReadConfig()
	codeType := viper.GetString("codeType")
	slug, diff := query.GetSlug(index)
	codeR, needR, titleR := query.GetInter(slug, codeType)
	data := &model.Data{Index: index, Diff: diff, CodeR: codeR, NeedR: needR, TitleR: titleR}
	codeDir := viper.GetString("codeDir")

	var dir string
	if viper.GetBool("moduled") {
		diff := strings.ToLower(data.Diff)
		dir = filepath.Join(codeDir, diff, strconv.Itoa(data.Index)+"."+data.TitleR)
	} else {
		dir = filepath.Join(codeDir, strconv.Itoa(data.Index)+"."+data.TitleR)
	}
	os.MkdirAll(dir, os.ModePerm)
	codeFilename := "main." + codeType
	codeFilePath := filepath.Join(dir, codeFilename)
	needFilePath := filepath.Join(dir, "need.md")
	if !isExist(codeFilePath) {
		color.Cyan("请勿重复拉取文件")
	} else {
		cfile, err := os.Create(codeFilePath)
		checkErr(err)
		defer cfile.Close()
		info := viper.GetString("append.prefix") + data.CodeR
		cfile.WriteString(info)
	}
	if viper.GetBool("function.useNeed") {
		if !isExist(needFilePath) {
			color.Cyan("请勿重复拉取文件")
		} else {
			nfile, err := os.Create(needFilePath)
			checkErr(err)
			defer nfile.Close()
			needR = convert(needR)
			nfile.WriteString(needR)
		}
	}
	t.Stop()
}

func convert(input string) (output string) {
	output = html2md.Convert(input)
	return
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return true
	}
	if os.IsExist(err) {
		return true
	}
	return false
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(getCmd)
}
