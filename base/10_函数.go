package main

import "fmt"

func add(a, b int) int {
	return a + b
}

// 同类型的参数可以合并
func sub(a, b int) int {
	return a - b
}

// 多返回值
// 第一个返回值是除法的结果，第二个是否成功
func div(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	} else {
		return a / b, true
	}
}

// 返回值也可以命令
func mul(a, b int) (result int) {
	// 带有命令的返回值会在进入函数的时候就进行初始化
	result = a * b
	return
}

// 不定参数
func addAll(a ...int) int {
	sum := 0
	for _, value := range a {
		sum += value
	}
	return sum
}

func main() {
	// 调用不定参数
	fmt.Print(addAll(1, 2, 3, 4))
	// 将一个[]int传入不定参数中
	nums := []int{1, 2, 3, 4}
	// nums...将参数变成一个一个的
	fmt.Print(addAll(nums...))
}
