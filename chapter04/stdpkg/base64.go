package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// base64.RawStdEncoding
	// base64.URLEncoding // 常用
	// base64.RawURLEncoding
	// base64.StdEncoding // 常用

	x := base64.StdEncoding.EncodeToString([]byte("i am guomianrui")) // 字节数组  base64编码
	fmt.Println(x)
	bytes, err := base64.StdEncoding.DecodeString(x)
	fmt.Println(string(bytes), err)
}
