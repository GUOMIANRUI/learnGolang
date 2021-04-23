package main

import (
	"fmt"
	"sort"
)

type User struct {
	ID    int
	Name  string
	Score float64
}

func main() {
	list := [][2]int{{1, 3}, {5, 9}, {4, 5}, {6, 2}, {5, 8}}

	// 按照数组元素的第二个元素进行排序
	sort.Slice(list, func(i, j int) bool {
		return list[i][1] < list[j][1] // 前面大的话会往后移
	})
	fmt.Println(list)

	// 结构体排序
	users := []User{{1, "kk", 6}, {2, "小灰", 5}, {3, "俗人", 9}, {4, "小小", 8}}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Score < users[j].Score
	})
	fmt.Println(users)
}
