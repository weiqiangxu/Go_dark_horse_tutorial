package main

import "fmt"

func myfunc(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data=", data)
	}
}

func test(args ...int) {
	//全部元素传递给myfunc
	//myfunc(args...)
	//只想把后两个参数传递给另外一个函数使用
	//myfunc(args[:2]...) //0-2但是不包括数字2，传递过去
	myfunc(args[2:]...) //从args[2]开始（包括本身），把后面的所有元素传递过去
}

func main() {
	test(1, 2, 3)
}
