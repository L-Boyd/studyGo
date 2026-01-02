package main

import (
	"fmt"
	"unsafe"
)

// 一些go基础内容
func main() {
	// 字符串可以用+连接
	fmt.Println("lby" + "tech")

	/*
		格式化字符串
		Sprintf会根据格式化参数生成字符串并返回
		PrintF会根据格式化参数生成字符串并输出
		%d 表示整型数字，%s 表示字符串
	*/
	var stockcode = 123
	var endDate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var targetUrl = fmt.Sprintf(url, stockcode, endDate)
	fmt.Println(targetUrl)
	fmt.Printf("The endDate is %s\n", endDate)

	/*
		声明变量的形式
		var 变量名 变量类型
		var 变量名1, 变量名2 变量类型
		var 变 量名 = 值（根据值自行判断类型）
		变量名 := 值
	*/
	var s1 string = "this is a string"
	var n1, n2 int = 1, 2
	var n3 = 3.1415
	n4 := 12345
	fmt.Println(s1, n1, n2, n3, n4)

	/*
		未初始的变量:
		bool为false
		数值类型为0
		字符串为空字符串
	*/
	var uninitV1 bool
	var uninitV2 int
	var uninitV3 string
	fmt.Println(uninitV1, uninitV2, uninitV3)

	/*
		常量用const标识
	*/
	const NAME = "LBY"
	fmt.Println(NAME)
	// 常量内函数必须是内置函数
	const (
		const1 = "abc"
		const2 = len(const1)
		const3 = unsafe.Sizeof(const1)
	)
	fmt.Println(const1, const2, const3)

	/*
		指针和地址
		&：返回变量存储地址
		*：这个变量是指针变量
	*/
	fmt.Println("=====指针和地址=====")
	var v1 = 4
	var ptr *int
	ptr = &v1
	fmt.Println(*ptr) // 4
	fmt.Println(ptr)  // 0xc00008c058
	v1 = 10
	fmt.Println(*ptr) // 10

	/*
		函数
		func function_name([parameter list]) [return_types] {
		    body
		}
	*/
	fmt.Println("=====函数=====")
	fmt.Println(f1(1, 2))
	var function_s1 = "s1"
	var function_s2 = "s2"
	function_s1, function_s2 = f2(function_s1, function_s2)
	fmt.Println(function_s1, function_s2)
}

func f1(num1, num2 int) int {
	return num1 + num2
}

// 多返回值
func f2(s1, s2 string) (string, string) {
	return s2, s1
}
