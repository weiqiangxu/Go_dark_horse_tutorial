// if else if的使用
package main

import "fmt"

func main() {
	a := 10
	if a == 10 {
		fmt.Println("a==10")
	} else { //else后面没有条件
		fmt.Println("a!=10")
	}
	// 2
	if a := 10; a == 10 {
		fmt.Println("a==10")
	} else {
		fmt.Println("a!=10")
	}

	a = 12
	if a == 10 {
		fmt.Println("a==10")
	} else if a > 10 {
		fmt.Println("a>10")
	} else {
		fmt.Println("===")
	}

}
