package main

import "encoding/json"

import "fmt"

func main(){
	// 创建一个map
	m := make(map[string]interface{},4)
	m["company"] = "itcast"
	m["subjects"] = []string{"go","c++"}
	m["isok"] = true
	m["price"] = 66.666

	// 编码生成json
	result,err := json.Marshal(m)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(result))
}