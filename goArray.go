package main

import "fmt"

/*
go的数组
var arrayName [arraySize]dataType
*/
func main() {
	var array1 = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("array1[%d] = %d\n", 3, array1[3])
}
