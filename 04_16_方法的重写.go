package main

import "fmt"

type Person struct{
	name string
	sex byte
	age int
}

// 有一个学生继承Person字段 成员和方法都继承了
type Student struct {
	Person //匿名字段
	id int
	addr string
}

// Person类型 实现了一个方法
func (tmp *Person) PrintInfo(){
	fmt.Printf("name=%s,sex=%c,age%d",tmp.name,tmp.sex,tmp.age)
}



// Student也实现了一个方法,这个方法和Person方法同名这种方法叫重写
func (tmp *Student) PrintInfo(){
	fmt.Println("Student:tmp = ",tmp)
}

func main() {
	s := Student{Person{"mike",'m',18},666,"bj"}
	// 就近原则：先找本作用域的方法，找不到再用继承的方法
	s.PrintInfo() //到底调用的是Person还是Student 结论是Student
	// 显示调用继承的方法
	s.Person.PrintInfo()

}