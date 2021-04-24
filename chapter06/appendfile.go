package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	path := "user.log"
	// 追加文件
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, os.ModePerm)
	fmt.Println(file, err)

	if err == nil {
		// file.WriteString(strconv.FormatInt(time.Now().Unix(), 10)) // strconv.FormatInt int64-->string
		// file.WriteString("\n")
		// file.Close()

		// 先不关文件 将内容写到日志里面
		log.SetOutput(file)
		log.Print(time.Now().Unix())
		file.Close()
	}
}
