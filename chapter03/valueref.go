package main

import "fmt"

func main() {
	// 值类型   引用类型
	// 将变量赋值给新的一个变量，并修改新变量的值，如果对旧变量 有影响 --> 引用类型   无影响 --> 值类型
	array := [3]string{"A", "B", "C"}
	slice := []string{"A", "B", "C"}

	arrayA := array
	sliceA := slice

	arrayA[0] = "Z"
	sliceA[0] = "Z"
	fmt.Println(arrayA, array) //[Z B C] [A B C]
	fmt.Println(sliceA, slice) //[Z B C] [Z B C]
	fmt.Printf("%p %p\n", &arrayA, &array)
	fmt.Printf("%p %p\n", sliceA, slice)

	// int bool float32 float64 array slice map 指针
	// 值类型：int bool float array 指针
	// 引用类型：slice map

	m := map[string]string{}
	mA := m
	mA["kk"] = "西安"
	fmt.Println(mA, m) // map[kk:西安] map[kk:西安]

	// 对于值类型来说 想要改变它的值类型 只能通过指针
	age := 30
	pointer := &age
	*pointer = 31
	fmt.Println(age, *pointer)
}
