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
	dirPath := filepath.Dir(path)

	fmt.Println(filepath.Base("c:/test/a.txt")) // a.txt
	fmt.Println(filepath.Base("c:/test/xxxx/")) // xxxx
	fmt.Println(filepath.Base(path))            //filepath.exe

	// 获取父目录
	fmt.Println(filepath.Dir("c:/test/a.txt")) // c:\test
	fmt.Println(filepath.Dir("c:/test/xxxx/")) // c:\test\xxxx
	fmt.Println(filepath.Dir(path))

	// 查看文件后缀
	fmt.Println(filepath.Ext("c:/test/a.txt"))
	fmt.Println(filepath.Ext(path))

	// 得到父目录后，一般用来拼写配置文件的地址
	iniPath := dirPath + "/conf/ip.ini"
	fmt.Println(iniPath) // 但是这样拼接不推荐 会出问题比如斜杠 C:\Users\ADMINI~1\AppData\Local\Temp\go-build177035946\b001\exe/conf/ip.ini
	//用join()函数拼接最好
	fmt.Println(filepath.Join(dirPath, "conf", "ip.ini")) // C:\Users\ADMINI~1\AppData\Local\Temp\go-build177035946\b001\exe\conf\ip.ini

	// fmt.Println(filepath.Glob("./*.exe"))
	fmt.Println(filepath.Glob("./[ab]*.exe"))  // 以ab开头的GO文件
	fmt.Println(filepath.Glob("./[ab]*/*.go")) // 当前目录下的ab开头目录中的GO文件

	filepath.Walk(".", func(path string, fileInfo os.FileInfo, err error) error {
		fmt.Println(path, fileInfo.Name())
		return nil
	}) // 遍历
}
