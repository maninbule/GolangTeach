package main

import "fmt"

func main() {
	// 定义一个int类型的变量,默认值为0
	var a int
	fmt.Println("a = ", a)
	// 定义一个int类型的变量,并初始化为1
	var b int = 1
	fmt.Println("b = ", b)
	// 使用自动推断来创建变量
	c := 2
	fmt.Println("c = ", c)
	// 一次性创建多个变量
	var d, e, f int = 1, 2, 3
	fmt.Println("d,e,f = ", d, e, f)

}
