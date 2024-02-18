package main

import (
	"fmt"
	"sort"
)

/*
如果想要将一个类型和一个排序规则进行绑定，就需要将排序函数和类型写在一起
*/

type point struct {
	x int
	y int
}

type pointArr []point

func (p pointArr) Len() int {
	return len(p)
}
func (p pointArr) Less(i, j int) bool {
	if p[i].x != p[j].x {
		return p[i].x < p[j].x
	} else {
		return p[i].y < p[j].y
	}
}
func (p pointArr) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	// 直接使用元素的话，需要在排序的时候再转换一下类型
	points := []point{
		{3, 4},
		{1, 2},
		{3, 1},
		{2, 5},
	}
	sort.Sort(pointArr(points))
	fmt.Println(points)

	// 直接使用type类型，就可以直接排序
	p2 := pointArr{
		{3, 4},
		{1, 2},
		{3, 1},
		{2, 5},
	}

	sort.Sort(p2)
	fmt.Println(p2)

	// 可以看到这样的方式其实有点麻烦，最好还是用匿名函数的方式
}
