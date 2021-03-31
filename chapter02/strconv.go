package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串  -->  其他类型
	// 字符串  -->  bool

	if v, err := strconv.ParseBool("trueeeee"); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err) // strconv.ParseBool: parsing "trueeeee": invalid syntax   "trueeeee"是错误的值
	}

	// 字符串  -->  int
	if v, err := strconv.Atoi("1023"); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	// 十六进制的 64位
	if v, err := strconv.ParseInt("64", 16, 64); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	// float类型
	if v, err := strconv.ParseFloat("1.1", 64); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	// 其他类型 -->  字符串
	sd := fmt.Sprintf("%d", 12)
	fmt.Println(sd)

	sf := fmt.Sprintf("%.3f", 12.01)
	fmt.Println(sf)

	fmt.Printf("%q\n", strconv.FormatBool(false))
	fmt.Printf("%q\n", strconv.Itoa(12))
	fmt.Printf("%q\n", strconv.FormatInt(12, 10))              // 十进制
	fmt.Printf("%q\n", strconv.FormatFloat(10.1, 'E', -1, 64)) // E 科学计数  -1保持原有精度
}
