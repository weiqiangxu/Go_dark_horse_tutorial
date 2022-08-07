package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	var p1 *Student = &Student{1, "mike", 'm', 18, "bj"}
	fmt.Println("p1=",p1)

	// 制定成员初始化 ,没有初始化的成员 自动赋值为0
	p2 := &Student{name:"mike",addr:"bj"}
	fmt.Printf("p2 type is %T\n",p2)
	fmt.Println("p2=",p2)
}
