// 类型别名 -- 非常重要

package main

import "fmt"

func main() {
	// 给int64起一个别名叫bigint
	type bigint int64
	var a bigint //等价var a int64
	fmt.Printf("type is %T\n", a)

	// 多个一起改别名
	type (
		long int64
		char byte
	)

	var b long = 11
	var ch char = 'a'
	fmt.Printf("b=%d,ch=%c\n", b, ch)

}
