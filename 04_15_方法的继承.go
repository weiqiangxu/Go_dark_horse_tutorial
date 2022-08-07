package main

import "fmt"

type Person struct{
	name string
	sex byte
	age int
}
// Person类型 实现了1个方法
func (tmp *Person)PrintInfo(){
	fmt.Print(tmp.name) //这里打印出的是name  string == mike 而不是指针
	fmt.Print(tmp) //这里打印的也不是指针，但是明明他是指针变量呀 为啥，看调用
}
// 有学生 继承Person字段 成员和方法都继承了
type Student struct {
	Person
	id int
	addr string
}
func main(){
	s := Student{Person{"mike",'m',18},666,"bj"}
	s.PrintInfo()

	(&s).PrintInfo() //此刻打印的是 mike  和 &{mike,..}
}