package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 标准输出 os.stdout
	fmt.Println("xxx")
	// 相当于：
	os.Stdout.Write([]byte("xxx"))

	bytes := make([]byte, 5)
	n, err := os.Stdin.Read(bytes)
	fmt.Println(n, err, string(bytes))

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())

	fmt.Fprintf(os.Stdout, "%T", 1) // 输出内容对象， 格式化类型， 内容    将1输出到终端
}
