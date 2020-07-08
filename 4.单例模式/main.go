package main

import (
	"fmt"
	"sync"
)

// Single 单例结构
type Single struct {
	data int
}

var singleton *Single
var once sync.Once // 内核信号，时时刻刻只能运行一个

// GetInterface 获取单例
func GetInterface() *Single {
	once.Do(func() {
		singleton = &Single{
			data: 11,
		}
	})
	return singleton
}

func main() {
	i1 := GetInterface()
	i2 := GetInterface()
	if i1 == i2 {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等")
	}
}
