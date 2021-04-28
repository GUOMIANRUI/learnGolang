package main

import (
	"fmt"
	"os"
	"strings"
)

// 获取文件夹的子文件信息
func main() {
	file, err := os.Open("xxx")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
		} else {
			file.Close()
		}
	}

	for _, path := range []string{"xxx", "reader.go", "test01"} {
		fileInfo, err := os.Stat(path) // 创建文件的时候可以用Stat这样就不用关闭文件了
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("文件不存在")
			}
		} else {
			fmt.Println(strings.Repeat("*", 20))
			fmt.Println(fileInfo.Name())
			fmt.Println(fileInfo.IsDir())
			fmt.Println(fileInfo.Size())    // 多少字节
			fmt.Println(fileInfo.ModTime()) // 最后修改时间

			if fileInfo.IsDir() {
				dirfile, err := os.Open(path)
				if err == nil {
					defer dirfile.Close()

					// childrens, _ := dirfile.Readdir(-1)
					// for _, children := range childrens {
					// 	fmt.Println(children.Name(), children.IsDir(), children.Size(), children.ModTime())
					// } // 文件都是有指针的 这里注释掉了下面读文件夹名才能读到

					names, _ := dirfile.Readdirnames(-1)
					for _, name := range names {
						fmt.Println(name)
					}
				}
			}
		}
	}

}
