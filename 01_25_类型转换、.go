// 类型转换

package main

import "fmt"

func main() {
	// 这种不能转换的类型,叫不兼容类型
	var flag bool
	flag = true
	// fmt.Printf("%d", flag) //%!d(bool=true) ----- error

	// bool不能转换为int
	// fmt.Printf("%d", int(flag)) //cannot convert flag (type bool) to type int

	// 0就是假，非0就是真
	// 整形也不能转换为 bool
	// flag = bool(1)
	fmt.Println(flag)
	var ch byte
	ch = 'a' //字符类型本质上就是整型
	var t int
	t = int(ch) //类型转换，把ch的值取出来后，转成int再给t赋值
	fmt.Println("t=", t)
}
