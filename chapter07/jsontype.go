package main

import (
	"encoding/json"
	"fmt"
)

const (
	Large  = iota //0 large   iota是golang语言的常量计数器, 出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次
	Medium        //1 medium
	Small         //2 small
)

type Size int

// 自定义时要注意定义  MarshalText() 和 UnmarshalText()
func (s Size) MarshalText() ([]byte, error) {
	switch s {
	case Large:
		return []byte("large"), nil
	case Medium:
		return []byte("medium"), nil
	case Small:
		return []byte("small"), nil
	default:
		return []byte("unknow"), nil
	}
}

// 再定义一个反序列化的 反序列化值 是要更改那个值 所以要用指针
func (s *Size) UnmarshalText(bytes []byte) error {
	switch string(bytes) {
	case "large":
		*s = Large
	case "medium":
		*s = Medium
	case "small":
		*s = Small
	default:
		*s = Small
	}
	return nil
}

func main() {
	var size Size = Medium
	bytes, _ := json.Marshal(size)
	fmt.Println(string(bytes)) // "medium"

	var size02 Size
	json.Unmarshal(bytes, &size02)
	fmt.Println(size02) // 0

	sizes := []Size{Large, Large, Small, Medium}
	bytes, _ = json.Marshal(sizes)
	fmt.Println(string(bytes)) // ["large","large","small","medium"]

	var sizes02 []Size
	json.Unmarshal(bytes, &sizes02) // 将bytes反序列化到sizes02
	fmt.Println(sizes02)

}
