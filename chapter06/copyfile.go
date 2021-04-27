package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func copyfile(src, dest string) {
	srcfile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	} else {
		defer srcfile.Close()
		destfile, err := os.Create(dest)
		if err != nil {
			fmt.Println(err)
		} else {
			defer destfile.Close()

			// 带缓冲
			reader := bufio.NewReader(srcfile)
			writer := bufio.NewWriter(destfile)

			bytes := make([]byte, 1024*1024*10) // 10M

			for {
				n, err := reader.Read(bytes)
				if err != nil {
					if err != io.EOF {
						fmt.Println(err)
					}
					break
				}
				destfile.Write(bytes[:n])
				writer.Flush()
			}

		}
	}
}

func main() {
	src := flag.String("s", "", "src file")
	dest := flag.String("d", "", "dest file")
	help := flag.Bool("h", false, "help")

	// 写一个帮助
	flag.Usage = func() {
		fmt.Println(`
Usage: copyfile -s srcfile -d destfile
Options:
		`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help || *src == "" || *dest == "" {
		flag.Usage()
	} else {
		copyfile(*src, *dest) // 把src复制到dest
	}

	fmt.Printf("%T, %T, %v, %v\n", src, dest, *src, *dest) // go run .\copyfile.go -s xxx -d xxx  *string, *string, xxx, xxx
}
