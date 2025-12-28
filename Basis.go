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

}
