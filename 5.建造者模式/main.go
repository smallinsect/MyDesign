package main

import "log"

// Builder 建造者模式
type Builder interface {
	Part1()
	Part2()
	Part3()
}

// "1"+"2"+"3"="123"
// 1+2+3=6

type Director struct {
	builder Builder //建造者接口
}

// NewDirector 创建接口
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Makedata() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

func main() {
	builder1 := &StringBuilder{
		result: "",
	}
	dict := NewDirector(builder1)
	dict.Makedata()
	log.Println(builder1.GetResult())

	builder2 := &IntBuilder{
		result: 0,
	}
	dict = NewDirector(builder2)
	dict.Makedata()
	log.Println(builder2.GetResult())
}
