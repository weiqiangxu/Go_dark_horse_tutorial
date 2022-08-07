package main

import "encoding/json"

import "fmt"


// 成员变量名首字母必须大写
type IT struct {
	Company string `json:"company"`  //二次编码
	Subjects []string
	IsOk bool `json:"isok"` //此字段不会输出到屏幕
	Price float64 `json:"price"` //转换为字符串后输出
}

func main(){
	jsonBuf := `
	{
		"company":"itcast",
		"subjects":[
			"go",
			"c++"
		],
		"isok":true,
		"price":66.66
	}
	`
	var tmp IT//定义一个结构体变量
	err := json.Unmarshal([]byte(jsonBuf),&tmp)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("tmp =%+v\n",tmp)
	// 仅仅获取自己想要的字段
	type IT2 struct {
		Subjects []string `json:"subjects"` //二次编码
	}
	var tmp2 IT2//定义一个结构体变量
	err = json.Unmarshal([]byte(jsonBuf),&tmp2)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("tmp2 =%+v\n",tmp2)
}