package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// 把计算的部分整合到一个函数 后面需要可以调用
func md5reader(reader *bufio.Reader) string { // 加了个bufio
	hasher := md5.New()
	bytes := make([]byte, 1024*1024*10)

	for {
		n, err := reader.Read(bytes)
		if err != nil {
			if err != io.EOF {
				return ""
			}
			break
		} else {
			hasher.Write(bytes[:n])
		}
	}
	return fmt.Sprintf("%X", hasher.Sum(nil))
}

func md5file(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		// fmt.Println(err)
		return "", err
	} else {
		defer file.Close()
		return md5reader(bufio.NewReader(file)), nil
	}
	return ""
}

func md5str(txt string) (string, error) {
	return md5reader(bufio.NewReader(strings.NewReader(txt))), nil // 装饰器
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
		var err error
		if *path != "" {
			md5, err = md5file(*path)
		} else {
			md5, err = md5str(*txt)
		}
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(md5)
		}
	}
}
