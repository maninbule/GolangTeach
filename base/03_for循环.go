package main

import "fmt"

func main() {
	// 标准的for循环
	for i := 1; i <= 10; i++ {
		fmt.Println("i = ", i)
	}
	// 其他语言的while循环
	cnt := 0
	for cnt < 10 {
		fmt.Println("cnt = ", cnt)
		cnt++
	}
	// 死循环
	j := 0
	for {
		j++
		if j < 10 {
			fmt.Println("j = ", j)
		} else {
			break
		}
	}
}
