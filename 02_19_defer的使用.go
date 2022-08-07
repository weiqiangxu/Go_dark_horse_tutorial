package main

import "main"

func main(){
	// defer延迟调用 main函数结束前调用
	defer fmt.Println("bbb")
	fmt.Println("aaa")
}