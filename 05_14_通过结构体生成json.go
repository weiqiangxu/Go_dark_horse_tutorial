package main

import "encoding/json"

import "fmt"

// 成员变量名首字母必须大写
type IT struct {
	Company string
	Subjects []string
	IsOk bool
	Price float64
}

func main(){
	//定义一个结构体变量 同时初始化
	s := IT{"itcast",[]string{"Go","C++"},true,66.1}
	// 编码 根据内容生成json文本
	buf ,err := json.Marshal(s)
	// 格式化编码 
	// buf ,err := json.MarshalIndent(s,""," ")
	if err != nil {
		fmt.Println("err = ",err)
		return
	}
	fmt.Println("buff = ",string(buf))
}