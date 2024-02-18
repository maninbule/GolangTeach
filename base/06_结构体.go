package main

import "fmt"

type Position struct {
	x int
	y int
}

func main() {
	// 显式定义结构体
	var p Position
	p.x = 1
	p.y = 2
	fmt.Println("p = ", p)
	// 自动推断
	q := Position{3, 4}
	fmt.Println("q = ", q)
}
