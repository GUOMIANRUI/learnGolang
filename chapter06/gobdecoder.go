package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Tel      string
	Addr     string
}

// 将文件中的内容提取 gob的decode方法  反序列化

func main() {
	users := map[int]User{}

	file, err := os.Open("user.gob")
	if err == nil {
		defer file.Close()

		decoder := gob.NewDecoder(file)
		decoder.Decode(&users)

		fmt.Println("%#v", users)
	}
}
