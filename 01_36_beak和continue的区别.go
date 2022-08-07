package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0

	for { // for后面不写任何东西表示无条件限制，永远循环

		i++

		time.Sleep(time.Second) //演示1s

		if i == 5 {
			// break 跳过当前循环  如果嵌套多个循环 跳出最近的那个循环
			continue //跳过本次循环 下一次继续
		}
		fmt.Println("i = ", i)
	}
}
