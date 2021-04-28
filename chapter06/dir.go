package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Mkdir("test01", 0644)
	// fmt.Println(os.Mkdir("test01", 0644)) 文件夹已存在，再创建会报错
	// os.Rename("test01", "test02")
	// os.Remove("test02")

	fmt.Println(os.MkdirAll("test01/xxx", 0644))
	fmt.Println(os.RemoveAll("test01"))
}
