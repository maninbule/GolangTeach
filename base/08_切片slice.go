package main

import "fmt"

func main() {
	// 没有显式定义，只能用make语句来创建

	// 完整定义
	C := make([]int, 2, 5)
	fmt.Println("C = ", C)
	// 创建一个空切片
	A := make([]int, 0)
	fmt.Println(A)
	A = append(A, 1)
	fmt.Println(A)
	A = append(A, 2, 3, 4)
	fmt.Println(A)
	B := []int{5, 6, 7}
	A = append(A, B...)
	fmt.Println(A)

	// 二维切片
	// 创建二维切片
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	// 切片深拷贝copy
	D := make([]int, len(A))
	copy(D, A)
	fmt.Println("D = ", D)

	// 切片操作
	A1 := []int{1, 2, 3, 4, 5}
	// 去掉末尾
	A1 = A1[:len(A)-1]
	// 去掉开始
	A1 = A1[1:]
	fmt.Print(A1)
	// 遍历切片
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	// for range循环
	for index, value := range arr {
		fmt.Println("index = ", index, " value = ", value)
	}
}
