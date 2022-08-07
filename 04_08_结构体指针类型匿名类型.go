package main

import "fmt"


type Person struct {
	name string
	sex byte
	age int
}
type Student struct {
	*Person //指针类型 - 这个变量将会是一个地址，初始化时候只能赋予取地址符&obj
	id int 
	addr string
}

func main(){
	s1 := Student{&Person{"mike",'m',18},666,"bj"}
	fmt.Print(s1)

	// 先定义变量
	var s2 Student
	s2.Person = new(Person)//分配空间
	s2.name = "yoyo"
	fmt.Print(s2)
}
