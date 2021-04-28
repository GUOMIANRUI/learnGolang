package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Abs获取绝对路径
	fmt.Println(filepath.Abs("."))
	fmt.Println(os.Args) // 第一个参数是当前运行的程序
	path, _ := filepath.Abs(os.Args[0])
	fmt.Println(filepath.Base("c:/test/a.txt")) // a.txt
	fmt.Println(filepath.Base("c:/test/xxxx/")) // xxxx
	fmt.Println(filepath.Base(path))            //filepath.exe

	// 获取父目录
	fmt.Println(filepath.Dir("c:/test/a.txt")) // c:\test
	fmt.Println(filepath.Dir("c:/test/xxxx/")) // c:\test\xxxx
	fmt.Println(filepath.Dir(path))
	// 得到父目录后，一般用来拼写配置文件的地址
}
