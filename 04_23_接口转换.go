package main

import "fmt"

type Humaner interface{
	sayhi()
}

type Personer interface{
	Humaner //匿名字段 继承了sayhi()
	sing(lrc string)
}

type Student struct {
	name string
	id int 
}

//Student实现了sayhi
func (tmp *Student) sayhi(){
	fmt.Printf("Student[%s,%d] sayhi\n",tmp.name,tmp.id)
}
func (tmp *Student) sing(lrc string){
	fmt.Println("Student 在唱着:",lrc)
}
func main()  {
	// 超级可以转换为子集 反过来不可以
	var iPro Personer //超集
	iPro = &Student{"mike",666}
	iPro.sayhi()
	var i Humaner //子集
	// iPro = i //err
	i = iPro //可以，超集可以转换为子集
	i.sayhi()
}