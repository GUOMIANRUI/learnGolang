package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

/*
md5sum
-s "abc" 计算-s参数字母的md5
-f path 计算-f后面路径文件的md5值
*/

func md5file(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		hasher := md5.New()

		bytes := make([]byte, 1024*1024*10)

		for {
			n, err := file.Read(bytes)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			} else {
				hasher.Write(bytes[:n])
			}
		}
		return fmt.Sprintf("%X", hasher.Sum(nil))
	}
	return ""
}

func md5str(txt string) string {
	return fmt.Sprintf("%X", md5.Sum([]byte(txt)))
}

func main() {
	// 定义命令行参数
	txt := flag.String("s", "", "md5 txt")
	path := flag.String("f", "", "file path")
	help := flag.Bool("h", false, "help")

	flag.Usage = func() {
		fmt.Println(`
Usage: md5.exe [-s 123abc] [-f filepath]	
Options:	
		`)
		flag.Parse()
	}
	flag.Parse()

	if *help || *txt == "" && *path == "" {
		flag.Usage()
	} else {
		var md5 string
		if *path != "" {
			md5 = md5file(*path)
		} else {
			md5 = md5str(*txt)
		}
		fmt.Println(md5)
	}
}
