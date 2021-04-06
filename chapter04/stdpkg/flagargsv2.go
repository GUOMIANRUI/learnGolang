package main

import (
	"flag"
	"fmt"
)

// 命令行参数解析  不用绑定参数的写法
func main() {

	port := flag.Int("p", 22, "ssh port")
	host := flag.String("H", "127.0.0.1", "ssh host")
	verbor := flag.Bool("v", false, "detail log")
	help := flag.Bool("h", false, "help message")

	flag.Usage = func() {
		fmt.Println("usage: flagargs [-H 127.0.0.1] [-P 22] [-v]")
		flag.PrintDefaults()
	}

	flag.Parse()

	fmt.Printf("%T %T %T %T", port, host, verbor, help)     // 指针类型 *int *string *bool *bool
	fmt.Printf("%v %v %v %v", *port, *host, *verbor, *help) // 拿指针对应的值 即解引用 用*
}
