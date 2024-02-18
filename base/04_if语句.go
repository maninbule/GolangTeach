package main

import "fmt"

func main() {
	// 带初始化变量的if
	// 在if中申明的变量，作用域在if语句里面
	x := 1
	if y := x + 5; y < 5 {
		fmt.Println("y = x + 5 结果小于5")
	} else if z := x + y; z < 10 {
		fmt.Println("x + y < 10", x, y, z)
	}
}
