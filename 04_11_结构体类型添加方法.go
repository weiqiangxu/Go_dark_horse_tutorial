package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//  带有接收者的函数叫方法
func (tmp Person) PrintInfo(){
	fmt.Println("tmp = ",tmp)
}

// 给结构体指针挂载一个func必须使用结构体指针变量调用
func (p *Person)SetInfo(n string,s byte,a int)  {
	p.name = n
	p.sex = s
	p.age = a
}

func main(){
	p := Person{"mike",'m',18}
	p.PrintInfo()

	// 定义一个结构体变量
	var p2 Person
	(&p2).SetInfo("yoyo",'f',22) //取地址
	p2.PrintInfo()
}