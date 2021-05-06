package main

import (
	"encoding/json"
	"fmt"
)

// 1.若需要进行序列化或反序列化的属性必须公开（大写开头）
type User struct {
	ID   int    `json:"-"`          // 属性的标签 标签名:标签   这样ID就不会被序列化
	Name string `json:"name"`       // 这样标签就会变成name
	Sex  int    `json:"sex,string"` // "sex": "0"    `json:"sex,int,omitempty"`当0值的时候不进行序列化
	tel  string
	Addr string
}

// 标签名、标签 底层原理是反射 可以看json源码

func main() {
	user := User{1, "kk", 0, "211354", "广州"}
	bytes, _ := json.MarshalIndent(user, "", "\t")
	fmt.Println(string(bytes))

	// 反序列化
	var user02 User
	json.Unmarshal(bytes, &user02)
	fmt.Println(user02)
}
