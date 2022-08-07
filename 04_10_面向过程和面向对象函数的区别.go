package main

import "fmt"

// 面向过程
func AddT(a int,b int) int{
	return a+b
}

// 面向对象 给某一个类型绑定一个函数
type long int 

func (tmp long) AddB(other long) long {
	return tmp + other
}

func main(){
	var result int
	result = AddT(1,1)
	fmt.Print(result)

	// 面向对象调用
	var a long = 2
	r := a.AddB(3)
	fmt.Println("r=",r)

}