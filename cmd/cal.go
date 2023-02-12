package cmd

import (
	"io/fs"
	"os"
	"path/filepath"
	"strconv"

	"github.com/levinion/letc/config"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var calCmd = &cobra.Command{
	Short: "题数统计",
	Use:   "cal",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Run() {
	config.ReadConfig()
	if viper.GetBool("moduled") {
		FullCal()
	} else {
		Cal()
	}
}

func init() {
	rootCmd.AddCommand(calCmd)
}

func Cal() {
	rootDir := viper.GetString("codeDir")
	count := countDir(rootDir)
	color.Cyan("已完成题数: %v", count)
}

func FullCal() {
	rootDir := viper.GetString("codeDir")
	easyDir := filepath.Join(rootDir, viper.GetString("alias.easy"))
	middleDir := filepath.Join(rootDir, viper.GetString("alias.medium"))
	hardDir := filepath.Join(rootDir, viper.GetString("alias.hard"))
	countEasy := countDir(easyDir)
	countMedium := countDir(middleDir)
	countHard := countDir(hardDir)

	if viper.GetBool("style.tableMod") {
		TablePrint(countEasy, countMedium, countHard)
	} else {
		SimplePrint(countEasy, countMedium, countHard)
	}
}

func countDir(rootDir string) (count int) {
	filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if checkType(info.Name()) {
				count++
			} else if info.Name() == "todo" {
				count--
			}
		}
		return nil
	})
	if count < 0 {
		count = 0
	}
	return
}

func TablePrint(countEasy, countMedium, countHard int) {
	sum := strconv.Itoa(countEasy + countMedium + countHard)
	data := [][]string{
		{"简单", strconv.Itoa(countEasy), sum},
		{"中等", strconv.Itoa(countMedium), sum},
		{"困难", strconv.Itoa(countHard), sum},
	}
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"难度", "题数", "总题数"})

	table.SetHeaderColor(tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.FgRedColor})

	table.SetBorder(false)
	table.SetAutoMergeCells(true)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	table.Rich(data[0], []tablewriter.Colors{
		{tablewriter.FgGreenColor}, {tablewriter.FgCyanColor}, {tablewriter.FgHiMagentaColor},
	})
	table.Rich(data[1], []tablewriter.Colors{
		{tablewriter.FgYellowColor}, {tablewriter.FgCyanColor}, {tablewriter.FgHiMagentaColor},
	})
	table.Rich(data[2], []tablewriter.Colors{
		{tablewriter.FgRedColor}, {tablewriter.FgCyanColor}, {tablewriter.FgHiMagentaColor},
	})

	table.Render()
}

func SimplePrint(countEasy, countMedium, countHard int) {
	color.Cyan("已完成总题数: %v\n", countEasy+countMedium+countHard)
	color.Green("简单: %v\n", countEasy)
	color.Yellow("中等: %v\n", countMedium)
	color.Red("困难: %v\n", countHard)
}

func checkType(name string) bool {
	ext := "." + viper.GetString("codeType")
	return filepath.Ext(name) == ext
}
