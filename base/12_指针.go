package main

import "fmt"

func changeInt(x int) {
	x++
}
func changeIntByPointer(x *int) {
	*x++
}

type student struct {
	name  string
	score int
}

func changeStruct(stu student) {
	stu.score = 100
}

func changeStructByPointer(stu *student) {
	(*stu).score = 100
}

func main() {
	x := 1
	changeInt(x)
	fmt.Println("changeInt修改x之后的值： x = ", x)

	changeIntByPointer(&x)
	fmt.Println("changeIntByPointer修改x之后的值： x = ", x)

	stu := student{"leo", 90}
	changeStruct(stu)
	fmt.Println("changeStruct修改score之后的值： stu.score = ", stu.score)

	changeStructByPointer(&stu)
	fmt.Println("changeStructByPointer修改score之后的值： stu.score = ", stu.score)
}
