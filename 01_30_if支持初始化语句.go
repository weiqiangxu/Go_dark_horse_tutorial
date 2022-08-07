package main

import "fmt"

func main() {
	s := "屌丝"

	if s == "王思聪" {
		fmt.Printf("左手一个妹纸，右手一个大妈")
	}

	// if支持1个初始化语句，初始化语句和判断条件以分号分隔
	if a := 10; a == 10 {
		fmt.Printf("a==10")
	}

}
