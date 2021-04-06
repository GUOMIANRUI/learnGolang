package main

import (
	"flag"
	"fmt"
)

// 命令行参数解析
func main() {
	var host string
	var port int
	var verbor bool
	var help bool

	// 绑定命令行参数与变量关系
	flag.IntVar(&port, "p", 22, "ssh port") // 最后面是帮助信息
	flag.StringVar(&host, "H", "127.0.0.1", "ssh host")
	flag.BoolVar(&verbor, "v", false, "detail log")
	flag.BoolVar(&help, "h", false, "help message")

	// 也可以改Usage函数
	flag.Usage = func() {
		fmt.Println("usage: flagargs [-H 127.0.0.1] [-P 22] [-v]")
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()
	// fmt.Println(host, port, verbor) // p = 传入的     go run flagargs.go -p 6666
	if help {
		flag.Usage()
	} else {
		fmt.Println(host, port, verbor)
		fmt.Println(flag.Args()) // 获取传入参数中多出来的参数 本质是切片
	}
}
