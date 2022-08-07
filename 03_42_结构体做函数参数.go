package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func test(s Student){
	s.id = 666
	fmt.Println("test : ",s)
}

func testTwo(s *Student){
	s.id = 666
}
func main(){
	s := Student{1,"mike",'m',18,"bj"}

	test(s)
	fmt.Println("main : ",s)

	testTwo(&s) //地址传递 引用传递 形参可以改变实参

	fmt.Println("main after testTwo : ",s)
}