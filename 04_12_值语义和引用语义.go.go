package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

// 接受者为普通变量，非指针，值语义，一份拷贝
func (p Person) SetInfoValue(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("SetInfoValue &p=", &p)
}

func (p *Person) SetInfoPointer(n string,s byte,a int){
	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("SetInfoPointer p = ",p)
}

func main(){
	var s1 Person
	fmt.Println("&s1 = ",&s1) //打印地址

	// 引用语义 
	(&s1).SetInfoPointer("mike",'m',18)
	fmt.Println("s1=",s1) //打印内容
}
