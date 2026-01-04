package main

import "fmt"

/*
go切片（动态数组）
*/
func main() {
	/*
		切片定义
	*/
	numbers := []int{1, 2, 3, 4, 5}
	// 初始长度为5
	numbers2 := make([]int, 5)
	// 初始长度为5, 容量为7
	numbers3 := make([]int, 5, 7)

	/*
		len和cap函数
	*/
	printSlice(numbers)
	printSlice(numbers2)
	printSlice(numbers3)

	/*
		切片截取
	*/
	fmt.Println(numbers)
	// 索引[1，3)
	fmt.Println("numbers[1:3] ==", numbers[1:3])
	// 默认下限为0
	fmt.Println("numbers[:3] ==", numbers[:3])
	// 默认上限为len(s)
	fmt.Println("numbers[2:] ==", numbers[2:])

	/*
		append和copy函数
	*/
	var numbers4 []int
	printSlice(numbers4)
	numbers4 = append(numbers4, 1)
	printSlice(numbers4)
	// 同时添加多个元素
	numbers4 = append(numbers4, 2, 3, 4)
	numbers5 := make([]int, len(numbers4), 2*cap(numbers4))
	copy(numbers5, numbers4)
	/*
		超出容量会扩容为原来的两倍
	*/
	for i := 0; i < 10; i++ {
		numbers5 = append(numbers5, i)
	}
	printSlice(numbers5)
}

func printSlice(numbers []int) {
	fmt.Println("len:", len(numbers), "cap:", cap(numbers), "slice:", numbers)
}
