package main

import (
	"fmt"
	"io"
	"strings"
)

// 情景：从网络传过来一个流数据，你不想保存到硬盘，在内存中直接去处理 可以用reader一部分一部分从IO读进来   可以像操作流一样去操作字符串
func main() {
	reader := strings.NewReader("abcdef\n1234556")
	bytes := make([]byte, 3)

	for {
		n, err := reader.Read(bytes)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		} else {
			fmt.Println(n, bytes[:n])
		}
	}

	// 不同于bufio bufio 是相对于缓冲（不是内存也不是磁盘）的io 可以一行一行打印
	// scanner := bufio.NewScanner(reader)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	// bufio 缓冲     os.OpenFile 磁盘   strings.builder.strings.writer 内存

	var builder strings.Builder // 声明一个结构体 也可以builder := strings.Builder{} 但不太好看
	// builder提升了字符串的处理速度 （字符串是不可变的 每次处理都要重新申请一块内存 生成新的字符串 指定到新的变量上去 对于builder来说，不用每次去申请内存...）
	builder.Write([]byte("abc"))
	builder.WriteString("abcdef!@#")

	fmt.Println(builder.String())

}
