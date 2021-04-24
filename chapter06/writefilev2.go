package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	err := ioutil.WriteFile("user.txt", []byte("xxxxxxxxxxxx"), os.ModePerm) // 会返回是否有错误
	fmt.Println(err)
}
