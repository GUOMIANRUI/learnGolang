package main

import (
	"encoding/gob"
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

// 将内存中的内容持久化到文件中 gob的encoder方法  序列化
// 自定义或者是内置的一些类型都可以进行持久化
func main() {
	users := map[int]User{
		1: User{1, "小明", time.Now(), "123456789", "山东"},
		2: User{2, "小月", time.Now(), "321456789", "山西"},
		3: User{3, "小黑", time.Now(), "213456789", "湖北"},
	}

	file, err := os.Create("user.gob")
	if err == nil {
		defer file.Close()
		encoder := gob.NewEncoder(file)
		encoder.Encode(users)
	}
}
