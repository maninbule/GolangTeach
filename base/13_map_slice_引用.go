package main

import "fmt"

// 记住，map和slice本身就是指针(引用)，不存在副本机制，只能使用copy去拷贝

func changeSlice(arr []int) {
	for i := 0; i < len(arr); i++ {
		(arr)[i] *= 2
	}
}

func changeMap(mp map[int]string) {
	mp[100] = "one hundred"
	delete(mp, 1)
}

func main() {
	// 参数传slice
	A := []int{1, 2, 3, 4, 5, 6}
	//changeSlice(&A[1:4])这样会报语法错误，slice默认不让你取地址
	//因为他本身就是地址
	changeSlice(A[1:4])
	fmt.Println(A)

	// 参数传map
	mp := map[int]string{
		1: "one",
		2: "two",
	}
	changeMap(mp)
	fmt.Println(mp)
}
