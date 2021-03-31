package main

import "fmt"

func main() {
	heights := []int{10, 6, 7, 9, 5}

	for i := 0; i < len(heights)-1; i++ {
		fmt.Println("第%d次排序", heights)
		for j := 0; j < len(heights)-1; j++ {
			if heights[j] > heights[j+1] {
				tmp := heights[j]
				heights[j] = heights[j+1]
				heights[j+1] = tmp
			}
		}
	}
	fmt.Println(heights)
}
