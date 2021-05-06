package main

import (
	"encoding/json"
	"fmt"
)

// 把一个变量变成一个json字符串

func main() {
	/*
		json.Marshal  用来做 序列化（内存 -> 字符串/字节切片）
		json.Unmarshal 用来做 反序列化（字符串/字节切片 -> 内存）
	*/
	names := []string{"小白", "小康", "小米", "小时"}
	users := []map[string]string{{"name": "bob", "addr": "汕头"}, {"name": "mary", "addr": "杭州"}, {"name": "jack", "addr": "广州"}}

	// bytes, err := json.Marshal(names)
	bytes, err := json.MarshalIndent(names, "", "\t") //做格式化 前缀是"" 缩进\t
	if err == nil {
		fmt.Println(bytes)
		// fmt.Println(string(bytes)) //["小白","小康","小米","小时"]
		fmt.Println(string(bytes))
		// 	[
		// 		"小白",
		// 		"小康",
		// 		"小米",
		// 		"小时"
		// ]
	}

	var names02 []string
	err = json.Unmarshal(bytes, &names02)
	fmt.Println(err)
	fmt.Println(names02) // [小白 小康 小米 小时] 反序列列化成了内存中的切片

	bytes, err = json.Marshal(users)
	if err == nil {
		fmt.Println(bytes)
		fmt.Println(string(bytes))
	}

	var users02 []map[string]string
	err = json.Unmarshal(bytes, &users02)
	fmt.Println(err)     // nil
	fmt.Println(users02) // [map[addr:汕头 name:bob] map[addr:杭州 name:mary] map[addr:广州 name:jack]]

	// 验证json的格式是否正确
	fmt.Println(json.Valid([]byte("[]xx"))) // false
}
