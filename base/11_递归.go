package main

// 和其他语言一样的，没什么特别的
func fib(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fib(n-1)
	}
}

func main() {

}
