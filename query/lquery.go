package query

import (
	_ "embed"
	"fmt"

	"github.com/levinion/letc/model"

	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

//go:embed resources/slugPayload.json
var slugPayload string

//go:embed resources/interPayload.json
var interPayload string

func GetSlug(index int) (slug, diff string) {
	res := query(slugPayload)
	jsonString := res.String()
	baseFindPath := fmt.Sprintf("data.allQuestionsBeta.#(questionFrontendId=%d)", index)
	slugFindPath := baseFindPath + ".titleSlug"
	diffFindPath := baseFindPath + ".difficulty"
	slug = gjson.Get(jsonString, slugFindPath).String()
	diff = gjson.Get(jsonString, diffFindPath).String()
	return
}

func GetInter(slug, codeType string) (string, string, string) {
	interPayload, err := sjson.Set(interPayload, "variables.titleSlug", slug)
	checkErr(err)
	res := query(interPayload)

	jsonString := res.String()
	codeTypeCode := model.CodeTypeMap[codeType]
	code := fmt.Sprintf("data.question.codeSnippets.%v.code", codeTypeCode)
	need := "data.question.translatedContent"
	title := "data.question.translatedTitle"
	codeR := gjson.Get(jsonString, code).String()
	needR := gjson.Get(jsonString, need).String() //html
	titleR := gjson.Get(jsonString, title).String()
	return codeR, needR, titleR
}

func query(bodyJsonString string) *req.Response {
	res, err := req.C().R().
		//csrf标头设置
		SetHeader("x-csrftoken",
			"RuQr5eiwJaQHSulKTOfVSSIgSOtJjEm7O9oPaEEkbFT27d16fmtWlp6byCWQiVT1").
		SetHeader("referer", "https://leetcode.cn/problems").
		SetBodyJsonString(bodyJsonString).
		Post("https://leetcode.cn/graphql/")
	checkErr(err)
	return res
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
