package main

import (
	"calc"
	"fmt"
)

func main(){
	a := calc.Add(10,20)
	fmt.Println("a=",a)

	fmt.Println("r=",calc.Minus(10,5))
}
// ./src/calc

// calc.go

// package calc

// 1 不同目录报名不一样
// 2 调用不同包里面的函数，格式：报名，函数名()
// 3 调用别的包的函数,这个包函数名称如果首字母是小写,无法让别人调
// 用，要想别人能调用必须首字母大写

