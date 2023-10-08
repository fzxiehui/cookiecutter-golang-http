package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

type Param struct {
	Name      string
	LowerName string
}

func createHandler(p Param) {
	handler, err := template.New("handler.tpl").Delims("{[", "]}").ParseFiles("./tpl/create/handler.tpl")
	if err != nil {
		panic(err)
	}
	filename := fmt.Sprintf("./internal/handler/%s.go", p.LowerName)
	fh, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	handler.Execute(fh, p)
}

func createService(p Param) {
	service, err := template.New("service.tpl").Delims("{[", "]}").ParseFiles("./tpl/create/service.tpl")
	if err != nil {
		panic(err)
	}
	filename := fmt.Sprintf("./internal/service/%s.go", p.LowerName)
	fh, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	service.Execute(fh, p)
}

func createRepository(p Param) {
	repository, err := template.New("repository.tpl").Delims("{[", "]}").ParseFiles("./tpl/create/repository.tpl")
	if err != nil {
		panic(err)
	}
	filename := fmt.Sprintf("./internal/repository/%s.go", p.LowerName)
	fh, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	repository.Execute(fh, p)
}

func createModel(p Param) {
	model, err := template.New("model.tpl").Delims("{[", "]}").ParseFiles("./tpl/create/model.tpl")
	if err != nil {
		panic(err)
	}
	filename := fmt.Sprintf("./internal/model/%s.go", p.LowerName)
	fh, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	model.Execute(fh, p)
}

func createRequest(p Param) {
	request, err := template.New("request.tpl").Delims("{[", "]}").ParseFiles("./tpl/create/request.tpl")
	if err != nil {
		panic(err)
	}
	filename := fmt.Sprintf("./internal/pkg/request/%s.go", p.LowerName)
	fh, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	request.Execute(fh, p)
}

func main() {
	// 接收参数
	args := os.Args
	// fmt.Print(args)
	if len(args) != 3 {
		fmt.Println("参数错误")
		return
	}

	model := args[2]
	p := Param{
		Name:      FirstUpper(model),
		LowerName: model,
	}
	f := args[1]
	if f == "all" || f == "handler" {
		createHandler(p)
	}
	if f == "all" || f == "service" {
		createService(p)
	}
	if f == "all" || f == "repository" {
		createRepository(p)
	}
	if f == "all" || f == "model" {
		createModel(p)
	}
	if f == "all" || f == "request" {
		createRequest(p)
	}

}
