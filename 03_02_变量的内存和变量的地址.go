package main

import "fmt"

func main() {
	var a int = 10
	// 每个变量有2层含义：变量的内存，变量的地址
	fmt.Printf("a=%d\n",a) //变量的内存
	fmt.Printf("&a=%v\n",&a)
}
