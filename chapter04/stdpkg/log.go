package main

import "log"

func main() {
	log.Printf("Printf日志：%s", "nin")
	log.SetPrefix("prefix:") // 给之后的日志设置前缀
	log.Printf("Printf日志：%s", "nin")

	// log.SetFlags(log.Flags() | log.Lshortfile) // 后面打印时会多一个文件名
	log.SetFlags(log.Flags() | log.Llongfile) // 后面打印时会多一个文件路径  prefix:2021/04/07 11:18:39 D:/Linux运维/GO/test_go/chapter04/stdpkg/log.go:14: Fatalf日志：nin

	// log.Panicf("Panic日志：%s", "nin") // 会报错 因为没有对panic恢复  不常用
	log.Fatalf("Fatalf日志：%s", "nin") // 打印完退出

}
