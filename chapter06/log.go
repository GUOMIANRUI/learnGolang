package main

import (
	"log"
	"os"
	"time"
)

func main() {
	path := "user.log"
	// 追加文件
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err == nil {
		// 将内容追加写到日志user.log里面 以后写日志可以这样写
		log.SetOutput(file)
		log.SetPrefix("users:") // 写个前缀
		log.SetFlags(log.Flags() | log.Lshortfile)
		log.Print(time.Now().Unix())
		file.Close()
	}
}
