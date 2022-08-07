package main

import "fmt"

type Person struct{
	name string
	sex byte
	age int
}

func (p Person)SetInfoValue(){
	fmt.Print("SetInfoValue:",p)
	fmt.Println("SetInfoValue")
}

// 这个虽然是挂载在结构体指针，但是直接结构体变量也是可以调用的
// 并且会自动转为引用传值
func (p *Person)SetInfoPointer(){
	fmt.Print("SetInfoPointer:",&p)
	p.age = 99 //会更新main的p结构体
	fmt.Println("SetInfoPointer")
}
func main(){
	p := Person{"mike",'m',18}
	fmt.Print(p)
	p.SetInfoPointer() 
	// 内部 先把p 转为&p再调用 (&p).SetInfoPointer
	p.SetInfoValue()
	fmt.Print("main ",p)
}	