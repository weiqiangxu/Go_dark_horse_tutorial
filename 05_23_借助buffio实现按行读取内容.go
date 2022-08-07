package main

import (
	"bufio"
	"io"
	"fmt"
	"os"
)

func ReadFileLine(path string) {
	// 打开文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	// 关闭文件
	defer f.Close()

	// 新建一个缓冲区 把内容放在缓冲区
	r := bufio.NewReader(f)

	for {
		// 遇到'\n'结束读取
		buf ,err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { //文件已经结束
				break
			}
			fmt.Println("err = ",err)
		}
		fmt.Printf("buf = %s\n",string(buf))
	}

}

func main() {
	path := "./demo.txt"
	ReadFileLine(path)
}
