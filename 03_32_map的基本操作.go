package main

import "fmt"

func main(){
	// 对于map只有len 没有cap 

	// 可以通过make创建
	// 可以指定长度,但是只是指定了容量，但是里面确实一个数据也没有
	m3 := make(map[int]string,2)
	m3[1] = "mike"
	fmt.Println("len = ",len(m3)) //注意这里输出不是2是1

	// 键值必须是唯一的   并且是无序的

	// 赋值
	m4 := map[int]string{1:"mike",2:"yoyo"}
	// 如果已经存在的key值 - 修改内容
	m4[1] = "c++"
	// 不存在的key值是扩容  和append类似
	m4[3] = "go"
	fmt.Println("m4 = ",m4)

	// map的遍历 - 遍历结果是无序的
	for k,v := range m4 {
		fmt.Println(k,v)
	}
	// 如何判断一个key值是否存在
	// 第一个返回值为key所对应的value，第二个返回值为key是否存在的条件 存在ok为true
	value,ok := m4[6]
	if ok == true {
		fmt.Print(value)
	} else {
		fmt.Println("不存在")
	}

	// 删除key为1的内容
	delete(m4,1)
	fmt.Print(m4)

	// map做函数参数是引用传递的
	test(m4)
	fmt.Print(m4)
}

// 引用传递
func test(m map[int]string){
	delete(m,2)
}