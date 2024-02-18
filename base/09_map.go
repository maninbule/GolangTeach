package main

import (
	"fmt"
)

func main() {
	// 只能使用make来创建
	mp := make(map[int]string)
	// 添加key-value
	mp[1] = "one"
	mp[2] = "two"
	mp[3] = "three"
	// 根据key取出value
	fmt.Println(mp[1])
	// 检查key是否有对应的value
	s, ok := mp[2]
	if !ok {
		fmt.Println("不存在key = 2")
	} else {
		fmt.Println("存在key = 2,其value = ", s)
	}
	// 删除key-value
	delete(mp, 2)
	s1, ok := mp[2]
	if !ok {
		fmt.Println("不存在key = 2")
	} else {
		fmt.Println("存在key = 2,其value = ", s1)
	}
	// 遍历map
	for key, value := range mp {
		fmt.Println(key, value)
	}
}
