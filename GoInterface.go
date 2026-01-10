package main

import (
	"fmt"
	"math"
)

// 定义接口：形状
type Shape interface {
	// 面积
	Area() float64
	// 周长
	Perimeter() float64
}

// 定义一个结构体：园
type Circle struct {
	// 半径
	Radius float64
}

// Circle 实现 Shape 接口（实现接口的两个方法就被认为实现了这个接口）
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	c := Circle{Radius: 5}
	var s Shape = c // 接口变量可以存储实现了接口的类型
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}
