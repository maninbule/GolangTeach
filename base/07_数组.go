package main

import "fmt"

func main() {
	// 显式定义数组
	var A [3]int
	A[0] = 1
	A[1] = 2
	A[2] = 3
	fmt.Println(A)

	// 自动推断
	B := [3]int{4, 5, 6}
	C := [3]int{} // 0 0 0
	fmt.Println("B = ", B)
	fmt.Println("C = ", C)

}
