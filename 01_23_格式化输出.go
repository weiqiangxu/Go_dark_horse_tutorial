// 格式化输出 - 不需要强行记忆  很容易忘记

// %s  %t

package main

import "fmt"

func main() {
	a := 97
	b := "abc"
	c := 'a'
	d := 3.14
	// %f 浮点型
	// %c 字符个数 - 可以吧输入的数字按照ASCII码转换为对应的字符
	// %d 整形格式
	// %s 字符串格式
	// %T操作变量所属类型
	// %v:自动匹配格式输出但不见得很智能 比如c输出了数字
	fmt.Printf("%T，%T，%T，%T\n", a, b, c, d)
	fmt.Printf("%v\n", c)
	fmt.Printf("%c", a)
}
