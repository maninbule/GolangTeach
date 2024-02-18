package main

import (
	"fmt"
	"os"
)

func main() {
	// 专门打印错误：打印出来字是红色的 (打印中间信息的时候比较好用)
	print("hello\n")

	// fmt打印
	// 打印不换行
	fmt.Print("hello")
	// 打印并换行
	fmt.Println("hello")
	// 格式化打印
	fmt.Printf("%s", "hello")
	// 打印成字符串
	s := fmt.Sprintf("%s", "test")
	fmt.Println(s)
	// 打印到文件
	fmt.Fprintf(os.Stderr, "hello_stderr")
	fmt.Fprintf(os.Stdout, "hello_stdout")
	file, err := os.OpenFile("base/in.txt", os.O_WRONLY, 0666)
	if err != nil {
		print("打开文件错误")
	}
	fmt.Fprintf(file, "输出到文件的内容")
	file.Close()

	// %v打印
	a := 1
	b := "string"
	c := [3]int{1, 2, 3}
	d := []int{4, 5, 6}
	e := make(map[int]string)
	e[1] = "one"
	e[2] = "two"

	fmt.Printf("%v %v %v %v %v\n", a, b, c, d, e)

}
