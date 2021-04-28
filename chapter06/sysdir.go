package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.TempDir())
	fmt.Println(os.UserCacheDir())
	fmt.Println(os.UserHomeDir())
	// 了解即可

	fmt.Println(os.Getwd()) // 运行路径（在哪运行exe） 一定一定不要修改它
}
