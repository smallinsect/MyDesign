package main

import "fmt"

// API 接口
type API interface {
	Say(name string) string
}

// Chinese 中国人
type Chinese struct{}

// Say 说中国话
func (*Chinese) Say(name string) string {
	return "你好，" + name
}

// English 英国人
type English struct{}

// Say 说英语
func (*English) Say(name string) string {
	return "hello，" + name
}

// NewAPI 简单工厂
func NewAPI(str string) API {
	if str == "en" {
		return &English{}
	}
	if str == "cn" {
		return &Chinese{}
	}
	return nil
}

func main() {
	api := NewAPI("en")
	fmt.Println(api.Say("pig"))

	api = NewAPI("cn")
	fmt.Println(api.Say("猪"))
}
