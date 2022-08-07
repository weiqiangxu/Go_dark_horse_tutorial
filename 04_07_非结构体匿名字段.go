package main

import "fmt"

type mystr string

type Person struct {
	name string
	sex byte
	age int
}
type Student struct {
	Person //结构体匿名字段
	int //基础类型的匿名字段
	mystr
}

func main(){
	s := Student{Person{"mike",'m',18},666,"hehehe"}
	fmt.Printf("s = %+v\n",s)

	s.Person = Person{"go",'m',22}
	fmt.Print(s)
	fmt.Println(s.int)

}