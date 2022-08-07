package main

import . "fmt" //调用函数，无需通过包名
import . "os"

// import io "fmt"//给报名起别名

func main(){
	Println("this is a test")
	Println("os.Args = ",Args)

	// io.Println("this is test two") //别名的调用
}

// 导入的包必须要使用否则编译不过

