package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 字节切片
	var bytes01 []byte = []byte{'a', 'b', 'c'}
	fmt.Printf("%T, %#v\n", bytes01, bytes01) // []uint8, []byte{0x61, 0x62, 0x63}

	s := string(bytes01) // 用strings转换
	fmt.Printf("%T %v\n", s, s)

	bs := []byte(s)
	fmt.Printf("%T %#v\n", bs, bs) // []uint8 []byte{0x61, 0x62, 0x63}

	// https://golang.google.cn/pkg/bytes/ 模块与strings的差不多
	fmt.Println(bytes.Compare([]byte("abc"), []byte("def")))
	fmt.Println(bytes.Index([]byte("abcdef"), []byte("def")))
	fmt.Println(bytes.Contains([]byte("abc"), []byte("a")))
}
