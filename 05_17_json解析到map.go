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
	// 创建一个map
	m := make(map[string]interface{},4)

	err := json.Unmarshal([]byte(jsonBuf),&m)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("m =%+v\n",m)
	var str string
	// 类型断言
	for key ,value := range m {
		switch data := value.(type) {
			case string:
				str = data
				fmt.Printf("map[%s]的值类型为string value=%s\n",key,str)
			case bool:
				fmt.Printf("map[%s]bool value=%s\n",key,data)
			case float64:
				fmt.Printf("map[%s]float64 value=%s\n",key,data)
			case []string:
				fmt.Printf("map[%s][]string value=%s\n",key,data)
			case []interface{}:
				fmt.Printf("map[%s][]string value=%s\n",key,data)
		}
	}
}