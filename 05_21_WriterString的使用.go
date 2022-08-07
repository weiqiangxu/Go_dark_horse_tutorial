package main

import "os"

import "fmt"

import "io"

func WriteFile(path string){
	// 打开文件 新建文件
	f,err := os.Create(path)
	if err != nil {
		fmt.Println("err = ",err)
		return
	}
	// 使用完毕 需要关闭文件
	defer f.Close()

	var buf string
	for i:=0;i<10;i++{
		// "i = 1\n" 这个字符串存储在buf中
		buf = fmt.Sprintf("i = %d\n",i)
		fmt.Println("buf = ",buf)

		n,err := f.WriteString(buf)
		if err!=nil {
			fmt.Println("err = ",err)
		}
		fmt.Println("n = ",n)
	}
}

func ReadFile(path string){
	// 打开文件
	f,err := os.Create(path)
	if err != nil {
		fmt.Println("err = ",err)
		return
	}
	// 关闭文件
	defer f.Close()

	buf := make([]byte,1024*2) //2k大小

	// n代表从文件读取内容的长度
	n , err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF { //文件出错 同时没有到结尾
		fmt.Println("err1 = ",err1)
		return
	}
	fmt.Println("buf = ",string(buf[:n]))

}

func test() {
	path := "./demo.txt"
	WriteFile(path)
}

func main(){
	path := "./demo.txt"
	ReadFile(path)
}