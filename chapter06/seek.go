package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("user.txt")

	bytes := make([]byte, 200)
	n, _ := file.Read(bytes)

	fmt.Println(n, bytes[:n])

	//seek()有两个参数： 偏移量 相对位置(文件首0 当前位置1 文件末尾2)
	//-5 向前偏移 5 向后偏移
	//文件首0 os.SEEK_SET
	//当前位置1 os.SEEK_CUR
	//文件末尾2 os.SEEK_END
	fmt.Println(file.Seek(0, 0)) // 移到文件的开始  0 <nil>

	n, err := file.Read(bytes)
	// 上一次已经把文件读完了 所以再读会有错误 0 EOF
	// 但是把文件指针放到开始，又可以读到内容了

	fmt.Println(n, err, bytes[:n])

	file.Close()
}
