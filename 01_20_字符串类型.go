// 字符串类型

package main

import "fmt"

func main() {
	var str1 string //声明变量
	str1 = "abc"
	fmt.Println(str1)

	// 自动推导类型
	str2 := "mike"
	fmt.Printf("%T", str2)

	// 内建函数 len
	fmt.Println(len(str2))

	// 字符和字符串的区别
	// 字符 -- 单引号 -- 往往都只有1个字符，转义字符除外'\n'
	var str string
	var ch byte
	ch = 'a'
	fmt.Println(ch)
	// 1、字符串 -- 双引号
	// 2、字符串有1个或者多个字符组成
	// 3、字符串都是隐藏了一个结束符，'\0' - 反斜杠零
	str = "a" //这个是由'a'和'\0'组成了一个字符串
	fmt.Println(str)

	str = "hello go"
	// 只想操作字符串的某一个字符 从0开始操作
	fmt.Println("str[0]=%c,str[1]=%c\n", str[0], str[1])
}
