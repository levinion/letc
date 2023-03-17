package query

import (
	"fmt"

	"github.com/levinion/gogpt"
	"github.com/spf13/viper"
)

func GetHelp(prompt string) string {
	if !viper.GetBool("ai.enable"){
		fmt.Println(viper.GetBool("ai.enable"))
		return "未开启ai功能，请勿使用此命令"
	}
	url := viper.GetString("ai.url")
	auth := viper.GetString("ai.auth")
	viper.SetDefault("max_tokens",100)
	if url == "" || auth == "" {
		return "请补全Header以开启ai功能"
	}
	c := gogpt.NewContext().
		SetHeader(&gogpt.Header{
			Url:  url,
			Auth: auth,
		}).
		SetBody(map[string]any{"max_tokens": viper.GetInt("ai.max_tokens")})
	c.Continue(prompt)
	return c.GetContent()
}
