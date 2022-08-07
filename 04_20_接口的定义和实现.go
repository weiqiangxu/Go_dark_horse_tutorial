package main

import "fmt"

type Humaner interface {
	// 方法 只有声明 没有实现 由别的类型（自定义类型）实现
	sayhi()
}

type Student struct {
	name string
	id int
}

// Student 实现了此方法
func (tmp *Student) sayhi(){
	fmt.Printf("Student [%s , %d] sayhi\n",tmp.name,tmp.id)
}

type Teacher struct {
	addr string
	group string
}

// Teacher实现了此方法
func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher [%s,%s] sayhi\n",tmp.addr,tmp.group)
}

type MyStr string

// MyStr实现了此方法

func (tmp *MyStr) sayhi(){
	fmt.Printf("MyStr[%s] sayhi\n",*tmp)
}

func main() {
	// 定义接口类型的常量
	var i Humaner

	// 只是实现了此接口方法的类型 那么这个类型的变量（接收者类型）就可以给i赋值
	s := &Student{"mike",666}
	i = s
	i.sayhi()

	t := &Teacher{"bj","go"}
	i = t
	i.sayhi()

	var str MyStr = "hellp mike"
	i = &str
	i.sayhi()
}


