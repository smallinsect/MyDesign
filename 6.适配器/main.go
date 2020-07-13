package main

import "fmt"

// 适配的目标接口
type Target interface {
	Request() string
}
type adapter struct {
	Adaptee
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

func (adap *adapter) Request() string {
	return adap.SpecficRequest()
}

//适配器
type Adaptee interface {
	SpecficRequest() string
}

// 载体，被适配的目标类
type adapteeImpl struct{}

// 工厂
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// 实际函数
func (ada *adapteeImpl) SpecficRequest() string {
	return "SpecficRequest小昆虫"
}

func main() {
	adapee := NewAdaptee()       // 适配器
	target := NewAdapter(adapee) // 传递进入
	res := target.Request()
	fmt.Println(res)
}
