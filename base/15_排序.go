package main

import (
	"fmt"
	"sort"
)

// 普通的排序，从小到大
// 适用的情况比较少
func sort_slice(arr []int) {
	sort.Ints(arr)
}

// 使用匿名函数排序,示例：从大到小排序
func sort_slice_by_cmp(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
}
func main() {
	arr := []int{3, 2, 1, 5, 3, 2}
	sort_slice(arr)
	fmt.Println("sort_slice 排序结果: ", arr)

	sort_slice_by_cmp(arr)
	fmt.Println("sort_slice_by_cmp 排序结果: ", arr)
}
