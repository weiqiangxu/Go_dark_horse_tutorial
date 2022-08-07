package main

import "fmt"

func main() {
	// 保存某一个变量的地址需要指针类型 *int 保存int的地址 **int保存*int地址
	// 声明（定义），定义只是特殊的声明
	// 定义一个变量p类型为*int
	var a int = 10
	var p *int
	p = &a //指针变量指向谁,就把谁的地址赋值给指针变量
	fmt.Printf("p =%v,&a=%v\n", p, &a)

	*p = 666 //*p操作的不是p的内存 是p所指向的内存(就是a)
	fmt.Printf("p =%v,&a=%v\n", *p, a)

	// p = 777 //会报错，因为指针变量只能传递指针不能赋值
}
