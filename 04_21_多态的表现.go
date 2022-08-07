package main

import "fmt"

type Humaner interface {
	// 方法 只有声明 没有实现 由别的类型（自定义类型）实现
	sayhi()
}

type Student struct {
	name string
	id   int
}

// Student 实现了此方法
func (tmp *Student) sayhi() {
	fmt.Printf("Student [%s , %d] sayhi\n", tmp.name, tmp.id)
}

type Teacher struct {
	addr  string
	group string
}

// Teacher实现了此方法
func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher [%s,%s] sayhi\n", tmp.addr, tmp.group)
}

type MyStr string

// MyStr实现了此方法

func (tmp *MyStr) sayhi() {
	fmt.Printf("MyStr[%s] sayhi\n", *tmp)
}

func WhoSayHi(i Humaner) {
	i.sayhi()
}

func main() {
	s := &Student{"mike", 666}
	t := &Teacher{"bj", "go"}
	var str MyStr = "hello mike"

	// 调用同一个函数 不同表现 多态 多种形态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	// 创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	// 第一个返回下标 第二个返回下标所对应的值
	for _, i := range x {
		i.sayhi()
	}
}
