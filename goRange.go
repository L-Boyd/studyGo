package main

import "fmt"

/*
go的range关键字
在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对
*/
func main() {
	/*
		数组和切片
	*/
	var nums = []int{1, 2, 4, 8, 16, 32}
	for i, v := range nums {
		fmt.Println(i, v)
	}

	/*
		字符串
	*/
	fmt.Println("======")
	var s = "hello"
	for i, c := range s {
		// 每个字符的索引和unicode码
		fmt.Println(i, c)
		fmt.Printf("%c\n", c)
	}
}
