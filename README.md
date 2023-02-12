# letc

![License](https://img.shields.io/badge/license-MIT-orange)
![Language](https://img.shields.io/badge/language-go-brightgreen)

[English](README_EN.md)

## 介绍
包含爬虫功能在内的力扣刷题辅助命令行工具

![演示](doc/show.jpg)

## 下载
```go
go install github.com/levinion/letc@latest
```

## 开始使用

运行 `letc init` 获取配置文件并开始使用

|命令|功能|
|----|----|
|`letc cal` 或 `letc`|统计已完成的题数,可在文件夹下新建名为 `todo` 的文件以忽略未做的题目|
|`letc get {题号}`|拉取给定题号的题目|

## 可配置选项

> 虽然此处未默认启用模块化，但推荐使用

```toml
codeDir="."         # 存放题目代码的主目录，默认为工作目录
moduled=false       # 是否启用模块化，若为true，则将题目按简单、中等、困难三个等级分类
codeType="go"       # 使用代码后缀区分语言类型

[function]
useNeed=true        # 是否在拉取题目的同时拉取需求，以MarkDown文件呈现

[append]
# 该字符串将在所拉取到的题目代码前插入
prefix="""\         
package main

func main(){

}

"""

[style]
tableMod=false      # 是否已表格模式显示完成题数，启用前需确保模块化开启

[alias]             # 此处提供自定义目录别名选项
easy="easy"
medium="medium"
hard="hard"
```